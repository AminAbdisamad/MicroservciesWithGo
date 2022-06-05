package handlers

import (
	"log"
	"net/http"

	"MicroservicesWithGo/models"
)

type Product struct {
	logger *log.Logger
}

func NewProduct(logger *log.Logger) *Product {
	return &Product{logger}
}

func (product *Product) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var products = models.GetProducts()
	err := products.ToJSON(writer)
	if err != nil {
		http.Error(writer, "Something went wrong", http.StatusBadRequest)
		product.logger.Println("something went wrong")
		return
	}

}
