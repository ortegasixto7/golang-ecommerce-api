package mongoDb

import (
	"context"
	"log"
	"os"

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
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
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

	db := client.Database(os.Getenv("DB_NAME"))
	ProductsCollection = db.Collection("products")
	AdminUsersCollection = db.Collection("adminUsers")
	UsersCollection = db.Collection("users")
}
