package fetcher

import (
	"math/rand"
	"sync"
	"time"
)

func FetchPricers(priceChannel chan<- float64) {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		priceChannel <- FetchPricersSite1()
	}()
	go func() {
		defer wg.Done()
		priceChannel <- FetchPricersSite2()
	}()
	go func() {
		defer wg.Done()
		priceChannel <- FetchPricersSite3()
	}()
	wg.Wait()
	close(priceChannel)
}

func FetchPricersSite1() float64 {
	time.Sleep(1 * time.Second)
	return rand.Float64() * 100
}

func FetchPricersSite2() float64 {
	time.Sleep(2 * time.Second)
	return rand.Float64() * 100
}

func FetchPricersSite3() float64 {
	time.Sleep(3 * time.Second)
	return rand.Float64() * 100
}
