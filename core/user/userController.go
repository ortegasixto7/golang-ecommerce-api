package user

import (
	"github.com/ortegasixto7/echo-go-supermarket-api/core/shared"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/user/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserPersistence IUserPersistence
}

func (this UserController) SignUp(request *requests.SignUpRequest) error {
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	user := User{
		Id:        primitive.NewObjectID().Hex(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Username:  request.Username,
		Password:  string(encryptedPassword),
		Role:      shared.USER,
	}
	this.UserPersistence.Save(user)
	return nil
}
