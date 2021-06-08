package model

import "time"

type Currency struct {
	Name           string    `json:"name" db:"name"`
	PriceInDollars float64   `json:"rate" db:"dollars_rate"`
	Date           time.Time `json:"date" db:"created_at"`
}
