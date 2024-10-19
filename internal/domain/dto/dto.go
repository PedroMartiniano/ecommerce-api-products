package dto

import "time"

type Stock struct {
	ID        string    `json:"id"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateProductStock struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Operation string `json:"operation"`
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryID  string  `json:"category_id"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
