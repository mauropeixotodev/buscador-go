package fetcher

import (
	"math/rand"
	"time"
)

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
