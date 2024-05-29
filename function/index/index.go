package index

import (
	"encoding/json"
	"fmt"

	"github.com/hsmtkk/vix-future-history/function/myhttp"
)

const URL = "https://serpapi.com/search.json"
const SERP_API_ENGINE = "google_finance"
const SERP_API_Q = "VIX:INDEXCBOE"

func Get(apiKey string) (float64, error) {
	query := map[string]string{
		"engine":  SERP_API_ENGINE,
		"q":       SERP_API_Q,
		"api_key": apiKey,
	}
	content, err := myhttp.Get(URL, query)
	if err != nil {
		return 0, err
	}
	fmt.Println(string(content))
	price, err := ParseJSON(content)
	if err != nil {
		return 0, err
	}
	return price, nil
}

type responseSchema struct {
	Markets struct {
		US []struct {
			Stock string  `json:"stock"`
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		} `json:"us"`
	} `json:"markets"`
}

func ParseJSON(content []byte) (float64, error) {
	rs := responseSchema{}
	if err := json.Unmarshal(content, &rs); err != nil {
		return 0, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	for _, record := range rs.Markets.US {
		if record.Stock == SERP_API_Q {
			return record.Price, nil
		}
	}
	return 0, fmt.Errorf("failed to find VIX data")
}
