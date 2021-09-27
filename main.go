package main

import (
	"fmt"

	"github.com/ortegasixto7/echo-go-supermarket-api/core/product"
	"github.com/ortegasixto7/echo-go-supermarket-api/persistence/mongoDb"
)

func main() {
	mongoDb.Setup()
	product := product.Product{
		Name:        "Pasta",
		Description: "Pasta Vermicelli",
		Price:       25.0,
		Quantity:    100,
	}
	fmt.Println("Hello")
	mongoDb.Save(product)

}
