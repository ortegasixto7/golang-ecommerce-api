package memory

import (
	"fmt"

	"github.com/ortegasixto7/echo-go-supermarket-api/core/product"
)

func Save(product *product.Product) {
	fmt.Println("Calling method save from in memory")
	fmt.Println(product)
}
