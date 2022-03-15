package auth

import (
	"github.com/golang-jwt/jwt"
)

type IAuthService interface {
	GenerateJwt(userId string) string
}

type AuthService struct{}

func (this AuthService) GenerateJwt(userId string) string {
	mySigningKey := []byte("AllYourBase")

	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Id:        userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(mySigningKey)
	return signedToken
}
