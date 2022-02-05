package mongoDb

import (
	"log"

	"github.com/ortegasixto7/echo-go-supermarket-api/core/product"
)

type MongoProductPersistence struct{}

func (this MongoProductPersistence) Save(product product.Product) {
	_, err := ProductsCollection.InsertOne(Ctx, product)
	if err != nil {
		log.Panicln(err)
	}
}

func (this MongoProductPersistence) Update(product product.Product) {

}

func (this MongoProductPersistence) Delete(id string) {

}

func (this MongoProductPersistence) GetAll() []product.Product {
	return []product.Product{}
}

func (this MongoProductPersistence) GetById(id string) product.Product {
	return product.Product{}
}
