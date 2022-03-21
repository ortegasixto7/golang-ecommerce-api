package auth

import (
	"errors"

	"github.com/ortegasixto7/echo-go-supermarket-api/core/admin"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/auth/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/auth"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/validations/customErrors"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	AuthService  auth.AuthService
	AdminService admin.AdminService
}

func (this AuthController) Login(request *requests.LoginRequest) (string, error) {
	adminUser, isEmpty := this.AdminService.GetByUserName(request.UserName)
	if isEmpty {
		return "", errors.New(customErrors.LOGIN_ERROR)
	}
	error := bcrypt.CompareHashAndPassword([]byte(adminUser.Password), []byte(request.Password))
	if error != nil {
		return "", errors.New(customErrors.LOGIN_ERROR)
	}
	return this.AuthService.GenerateJwt(adminUser.Id), nil
}

// TODO: Add middleware to verify requests
