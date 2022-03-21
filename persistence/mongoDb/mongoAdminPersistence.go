package mongoDb

import (
	"log"

	"github.com/ortegasixto7/echo-go-supermarket-api/core/admin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoAdminPersistence struct{}

func (this MongoAdminPersistence) Save(data admin.AdminUser) {
	_, err := AdminUsersCollection.InsertOne(Ctx, data)
	if err != nil {
		log.Panicln(err)
	}
}

func (this MongoAdminPersistence) Update(data admin.AdminUser) {
	filter := bson.M{"_id": data.Id}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "firstName", Value: data.FirstName},
			primitive.E{Key: "lastName", Value: data.LastName},
		}},
	}
	_, err := AdminUsersCollection.UpdateOne(Ctx, filter, update)
	if err != nil {
		log.Panicln(err)
	}
}

func (this MongoAdminPersistence) GetById(id string) admin.AdminUser {
	var data admin.AdminUser
	filter := bson.M{"_id": id}
	err := AdminUsersCollection.FindOne(Ctx, filter).Decode(&data)
	if err != nil {
		// fmt.Println(err)
	}
	return data
}

func (this MongoAdminPersistence) GetByUserName(username string) admin.AdminUser {
	var data admin.AdminUser
	filter := bson.M{"username": username}
	err := AdminUsersCollection.FindOne(Ctx, filter).Decode(&data)
	if err != nil {
		// fmt.Println(err)
	}
	return data
}
