package future

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"

	"github.com/hsmtkk/vix-future-history/function/myhttp"
)

const URL = "https://www.cboe.com/us/futures/market_statistics/settlement/csv"

type Data struct {
	Symbol string
	Expire time.Time
	Price  float64
}

func Get() ([]Data, error) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	dt := yesterday.Format("2006-01-02")
	query := map[string]string{
		"dt": dt,
	}
	content, err := myhttp.Get(URL, query)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(content))

	results := []Data{}

	return results, nil
}

type FilteredCSV struct {
	Product    string
	Symbol     string
	Expiration time.Time
	Price      float64
	MonthIndex int
}

func FilterCSV(parsedCSVs []ParsedCSV) []FilteredCSV {
	monthIndex := 1
	results := []FilteredCSV{}
	for _, parsedCSV := range parsedCSVs {
		if parsedCSV.Product != "VX" {
			continue
		}
		results = append(results, FilteredCSV{Product: parsedCSV.Product, Symbol: parsedCSV.Symbol, Expiration: parsedCSV.Expiration, Price: parsedCSV.Price, MonthIndex: monthIndex})
		monthIndex += 1
	}
	return results
}

type ParsedCSV struct {
	Product    string
	Symbol     string
	Expiration time.Time
	Price      float64
}

func ParseCSV(content []byte) ([]ParsedCSV, error) {
	results := []ParsedCSV{}
	reader := csv.NewReader(bytes.NewReader(content))
	reader.FieldsPerRecord = 4
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read as CSV: %w", err)
	}
	for _, record := range records[1:] {
		result, err := parseRecord(record)
		if err != nil {
			fmt.Println(err)
			continue
		}
		results = append(results, result)
	}
	return results, nil
}

func parseRecord(record []string) (ParsedCSV, error) {
	product := record[0]
	symbol := record[1]
	expirationStr := record[2]
	priceStr := record[3]
	expiration, err := time.Parse("2006-01-02", expirationStr)
	if err != nil {
		return ParsedCSV{}, fmt.Errorf("failed to parse %s as time: %w", expirationStr, err)
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return ParsedCSV{}, fmt.Errorf("failed to parse %s as float: %w", priceStr, err)
	}
	return ParsedCSV{Product: product, Symbol: symbol, Expiration: expiration, Price: price}, nil
}
