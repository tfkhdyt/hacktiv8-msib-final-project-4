package dto

import "time"

type ProductData struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Price     uint      `json:"price"`
	Stock     uint      `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
