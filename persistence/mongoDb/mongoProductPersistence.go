package mongoDb

import (
	"log"

	"github.com/ortegasixto7/echo-go-supermarket-api/core/product"
)

func Save(product product.Product) {
	result, err := ProductsCollection.InsertOne(Ctx, product)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(result.InsertedID)
}
