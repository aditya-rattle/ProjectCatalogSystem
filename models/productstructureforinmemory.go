package models

type InMemoryProduct struct {
	Id          int64 `json:"id"`
	Name        string `json:"name"`
	Price       int64 `json:"price"`
	Quantity    int64 `json:"quantity"`
	Description string `json:"description"`
}

