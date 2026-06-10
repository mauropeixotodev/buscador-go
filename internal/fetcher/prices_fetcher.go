package fetcher

import (
	"buscador/internal/models"
	"math/rand"
	"sync"
	"time"
)

func FetchPricers(priceChannel chan<- models.Price) {
	wg := sync.WaitGroup{}
	wg.Add(4)
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
	go func() {
		defer wg.Done()
		FetchAndSendMultiplePrices(priceChannel)
	}()
	wg.Wait()
	close(priceChannel)
}

func FetchPricersSite1() models.Price {
	time.Sleep(1 * time.Second)
	return models.Price{
		StoreName: "Site 1",
		Price:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchPricersSite2() models.Price {
	time.Sleep(2 * time.Second)
	return models.Price{
		StoreName: "Site 2",
		Price:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchPricersSite3() models.Price {
	time.Sleep(1 * time.Second)
	return models.Price{
		StoreName: "Site 1",
		Price:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchAndSendMultiplePrices(priceChannel chan<- models.Price) {
	time.Sleep(1 * time.Second)
	prices := []models.Price{
		models.Price{
			StoreName: "Site 4",
			Price:     rand.Float64() * 100,
			Timestamp: time.Now(),
		},
		models.Price{
			StoreName: "Site 5",
			Price:     rand.Float64() * 100,
			Timestamp: time.Now(),
		},
		models.Price{
			StoreName: "Site 6",
			Price:     rand.Float64() * 100,
			Timestamp: time.Now(),
		},
	}
	for _, price := range prices {
		priceChannel <- price
	}
}
