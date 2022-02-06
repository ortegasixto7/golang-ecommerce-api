package mongoDb

import (
	"log"

	"github.com/ortegasixto7/echo-go-supermarket-api/core/product"
	"go.mongodb.org/mongo-driver/bson"
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
	products := []product.Product{}
	cursor, err := ProductsCollection.Find(Ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(Ctx)
	for cursor.Next(Ctx) {
		var product product.Product
		if err = cursor.Decode(&product); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	return products
}

func (this MongoProductPersistence) GetById(id string) product.Product {
	return product.Product{}
}
