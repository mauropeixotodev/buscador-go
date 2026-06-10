package processor

import "fmt"

func ShowPrices(priceChannel <-chan float64, doneChannel chan<- bool) {
	var totalPrice float64
	var count int
	for price := range priceChannel {
		totalPrice += price
		count++
		fmt.Printf("Price: %f\n", price)
		fmt.Printf("Average price until now: %f\n", totalPrice/float64(count))
	}
	doneChannel <- true
}
