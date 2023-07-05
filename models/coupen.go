package models

import "time"

type Coupen struct {
	Code      string
	Expiry    time.Time
	MinAmount int
	Amount    int
	Status    string
}
