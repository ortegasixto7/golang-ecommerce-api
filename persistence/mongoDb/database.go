package mongoDb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	UsersCollection      *mongo.Collection
	ProductsCollection   *mongo.Collection
	AdminUsersCollection *mongo.Collection
	Ctx                  = context.TODO()
)

func Setup() {
	host := "127.0.0.1"
	port := "27017"
	connectionURI := "mongodb://" + host + ":" + port + "/"
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("goTest")
	ProductsCollection = db.Collection("products")
	AdminUsersCollection = db.Collection("adminUsers")
	UsersCollection = db.Collection("users")
}
