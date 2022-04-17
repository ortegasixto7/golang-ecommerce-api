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
	AuthService      auth.IAuthService
	AdminPersistence admin.IAdminPersistence
}

func (this AuthController) Login(request *requests.LoginRequest) (authToken string, error error) {
	adminUser, isEmpty := this.AdminPersistence.GetByUserName(request.UserName)
	if isEmpty {
		return "", errors.New(customErrors.LOGIN_ERROR)
	}
	hashError := bcrypt.CompareHashAndPassword([]byte(adminUser.Password), []byte(request.Password))
	if hashError != nil {
		return "", errors.New(customErrors.LOGIN_ERROR)
	}
	return this.AuthService.GenerateJwt(adminUser.Id), nil
}

func (this AuthController) SignUp() (authToken string, error error) {
	return "", nil
}
