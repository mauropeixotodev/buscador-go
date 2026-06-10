package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/models"
	"buscador/internal/processor"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan models.Price)
	doneChannel := make(chan bool)

	go fetcher.FetchPricers(priceChannel)
	go processor.ShowPrices(priceChannel, doneChannel)

	<-doneChannel

	fmt.Printf("Time taken: %s\n", time.Since(start))
	fmt.Printf("Time taken: %s\n", time.Since(start))
}
