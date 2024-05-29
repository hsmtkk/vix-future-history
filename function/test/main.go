package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hsmtkk/vix-future-history/function/future"
	"github.com/hsmtkk/vix-future-history/function/index"
)

func main() {
	futureData, err := future.New().Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(futureData)

	indexData, err := index.Get(os.Getenv("SERP_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(indexData)
}
