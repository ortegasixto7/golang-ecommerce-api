package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type IAuthService interface {
	GenerateJwt(userId string) (token string)
	DecodeJwt(token string) (userId string)
}

type AuthService struct{}

type Claims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func (this AuthService) GenerateJwt(userId string) (token string) {
	mySigningKey := []byte(os.Getenv("AUTH_SECRET_KEY"))

	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}

	result := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ = result.SignedString(mySigningKey)
	return token
}

func (this AuthService) DecodeJwt(token string) (userId string) {
	claims := &Claims{}
	_, error := jwt.ParseWithClaims(token, claims, func(decodedToken *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("AUTH_SECRET_KEY")), nil
	})

	if error != nil {
		return ""
	}

	return claims.UserId
}
