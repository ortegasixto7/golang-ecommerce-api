package mongoDb

import (
	"log"

	"github.com/ortegasixto7/echo-go-supermarket-api/core/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoUserPersistence struct{}

func (this MongoUserPersistence) Save(data user.User) {
	_, err := UsersCollection.InsertOne(Ctx, data)
	if err != nil {
		log.Panicln(err)
	}
}

func (this MongoUserPersistence) Update(data user.User) {
	filter := bson.M{"_id": data.Id}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "firstName", Value: data.FirstName},
			primitive.E{Key: "lastName", Value: data.LastName},
		}},
	}
	_, err := UsersCollection.UpdateOne(Ctx, filter, update)
	if err != nil {
		log.Panicln(err)
	}
}

func (this MongoUserPersistence) GetById(id string) (user.User, bool) {
	var data user.User
	filter := bson.M{"_id": id}
	err := UsersCollection.FindOne(Ctx, filter).Decode(&data)
	if err != nil {
		// fmt.Println(err)
	}
	return data, data.Id == ""
}

func (this MongoUserPersistence) GetByUserName(username string) (user.User, bool) {
	var data user.User
	filter := bson.M{"username": username}
	err := UsersCollection.FindOne(Ctx, filter).Decode(&data)
	if err != nil {
		// fmt.Println(err)
	}
	return data, data.Id == ""
}
