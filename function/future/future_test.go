package future_test

import (
	"os"
	"testing"

	"github.com/hsmtkk/vix-future-history/function/future"
	"github.com/stretchr/testify/assert"
)

func TestParseCSV(t *testing.T) {
	content, err := os.ReadFile("./sample.csv")
	assert.Nil(t, err)
	parsedCSV, err := future.ParseCSV(content)
	assert.Nil(t, err)
	assert.Equal(t, "VX", parsedCSV[0].Product)
	assert.Equal(t, "VX22/K4", parsedCSV[0].Symbol)
	assert.Equal(t, 13.125, parsedCSV[0].Price)
}
