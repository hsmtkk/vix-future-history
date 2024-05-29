package index_test

import (
	"os"
	"testing"

	"github.com/hsmtkk/vix-future-history/function/index"
	"github.com/stretchr/testify/assert"
)

func TestParseJSON(t *testing.T) {
	content, err := os.ReadFile("./sample.json")
	assert.Nil(t, err)
	price, err := index.ParseJSON(content)
	assert.Nil(t, err)
	assert.Equal(t, price, 12.92)
}
