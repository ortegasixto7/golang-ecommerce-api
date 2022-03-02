package mongoDb

import (
	"log"

	"github.com/ortegasixto7/echo-go-supermarket-api/core/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoProductPersistence struct{}

func (this MongoProductPersistence) Save(product product.Product) {
	_, err := ProductsCollection.InsertOne(Ctx, product)
	if err != nil {
		log.Panicln(err)
	}
}

func (this MongoProductPersistence) Update(product product.Product) {
	objectId, _ := primitive.ObjectIDFromHex(product.Id)
	filter := bson.M{"_id": objectId}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "name", Value: product.Name},
			primitive.E{Key: "description", Value: product.Description},
			primitive.E{Key: "price", Value: product.Price},
			primitive.E{Key: "quantity", Value: product.Quantity},
			primitive.E{Key: "categories", Value: product.Categories},
			primitive.E{Key: "photoUrl", Value: product.PhotoUrl},
		}},
	}
	_, err := ProductsCollection.UpdateOne(Ctx, filter, update)
	if err != nil {
		log.Panicln(err)
	}
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
	product := product.Product{}
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	err := ProductsCollection.FindOne(Ctx, filter).Decode(&product)
	if err != nil {
		log.Fatal(err)
	}
	return product
}
