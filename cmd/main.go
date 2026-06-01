package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/processor"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)

	wg, showWG := sync.WaitGroup{}, sync.WaitGroup{}
	wg.Add(3)
	showWG.Add(1)
	go func() {
		defer showWG.Done()
		processor.ShowPrices(priceChannel)
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPricersSite1()
	}()
	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPricersSite2()
	}()
	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPricersSite3()
	}()
	wg.Wait()
	close(priceChannel)
	showWG.Wait()

	fmt.Printf("Time taken: %s\n", time.Since(start))
	fmt.Printf("Time taken: %s\n", time.Since(start))
}
