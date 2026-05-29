package main

import (
	"buscador/internal/fetcher"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		var totalPrice float64
		var count int
		for price := range priceChannel {
			totalPrice += price
			count++
			fmt.Printf("Price: %f\n", price)
			fmt.Printf("Average price until now: %f\n", totalPrice/float64(count))
		}
		//fmt.Printf("Total price: %f\n", totalPrice)
		//fmt.Printf("Average price: %f\n", totalPrice/float64(count))
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

	fmt.Printf("Time taken: %s\n", time.Since(start))
	fmt.Printf("Time taken: %s\n", time.Since(start))
}
