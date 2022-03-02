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
	objectId, _ := primitive.ObjectIDFromHex(data.Id)
	filter := bson.M{"_id": objectId}
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
	data := admin.AdminUser{}
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	err := AdminUsersCollection.FindOne(Ctx, filter).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
