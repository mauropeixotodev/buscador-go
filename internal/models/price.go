package models

import "time"

type Price struct {
	StoreName string    `json:"store_name"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}
