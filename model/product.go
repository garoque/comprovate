package model

import "time"

type Product struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
