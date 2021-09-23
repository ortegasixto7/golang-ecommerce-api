package main

import (
	"fmt"

	"github.com/ortegasixto7/echo-go-supermarket-api/core/product"
	"github.com/ortegasixto7/echo-go-supermarket-api/persistence/memory"
)

func main() {
	product := product.Product{
		Id:          "guid",
		Name:        "Pasta",
		Description: "Pasta Vermicelli",
		Price:       25.0,
	}
	fmt.Println("Hello")
	memory.Save(&product)
}
