package mongoDb

import (
	"log"

	"github.com/ortegasixto7/echo-go-supermarket-api/core/product"
)

func SaveBackup(product product.Product) {
	result, err := ProductsCollection.InsertOne(Ctx, product)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(result.InsertedID)
}

type MongoProductPersistence struct{}

func (repository *MongoProductPersistence) Save(product product.Product) {
	log.Println(product)
}

func (repository *MongoProductPersistence) Update(product product.Product) {

}

func (repository *MongoProductPersistence) Delete(id string) {

}

func (repository *MongoProductPersistence) GetAll() []product.Product {
	return []product.Product{}
}

func (repository *MongoProductPersistence) GetById(id string) product.Product {
	return product.Product{}
}
