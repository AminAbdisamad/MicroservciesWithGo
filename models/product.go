package models

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SKU         string  `json:"SKU"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

type Products []Product

func (products *Products) ToJSON(writer io.Writer) error {
	var encoder = json.NewEncoder(writer)
	return encoder.Encode(products)
}

func GetProducts() Products {
	return productList
}

var productList = []Product{
	{
		ID:          1,
		Name:        "Mocha",
		Description: "White chocolate mocha",
		Price:       3.45,
		SKU:         "23A",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Late",
		Description: "Coffee Late",
		Price:       3.20,
		SKU:         "23B",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}
