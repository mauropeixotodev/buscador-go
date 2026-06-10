package processor

import (
	"buscador/internal/models"
	"fmt"
	"time"
)

func ShowPrices(priceChannel <-chan models.Price, doneChannel chan<- bool) {
	var totalPrice float64
	var count int
	for price := range priceChannel {
		totalPrice += price.Price
		count++
		fmt.Printf("Price: %f\n", price.Price)
		fmt.Printf("Store Name: %s\n", price.StoreName)
		fmt.Printf("Timestamp: %s\n", price.Timestamp.Format(time.RFC3339))
		fmt.Printf("Average price until now: %f\n", totalPrice/float64(count))
	}
	doneChannel <- true
}
