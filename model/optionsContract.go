package model

import "time"

// OptionsContract represents the data structure of an options contract
type OptionsContract struct {
	StrikePrice    float64   `json:"strike_price"`
	Type           string    `json:"type"`
	Bid            float64   `json:"bid"`
	Ask            float64   `json:"ask"`
	LongShort      string    `json:"long_short"`
	ExpirationDate time.Time `json:"expiration_date"`
}
