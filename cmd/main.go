package main

import (
	"buscador/internal/fetcher"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var price1, price2, price3 float64

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		price1 := fetcher.FetchPricersSite1()
		fmt.Printf("Price 1: %f\n", price1)

	}()
	go func() {
		defer wg.Done()
		price2 := fetcher.FetchPricersSite2()
		fmt.Printf("Price 2: %f\n", price2)
	}()
	go func() {
		defer wg.Done()
		price3 := fetcher.FetchPricersSite3()
		fmt.Printf("Price 3: %f\n", price3)
	}()
	wg.Wait()

	fmt.Printf("Time taken: %s\n", time.Since(start))

	fmt.Printf("Price 1: %f\n", price1)
	fmt.Printf("Price 2: %f\n", price2)
	fmt.Printf("Price 3: %f\n", price3)
	fmt.Printf("Time taken: %s\n", time.Since(start))
}
