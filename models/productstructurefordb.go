package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DbProduct struct {
	gorm.Model
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
	Description string `json:"description"`
}

type UserProduct struct {
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
	Description string `json:"description"`
}

type Transaction struct {
	Name      string    `json:"name"`
	Quantity  int64     `json:"quantity"`
	OrderTime time.Time `json:"ordertime"`
}
