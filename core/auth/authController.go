package auth

import (
	"github.com/ortegasixto7/echo-go-supermarket-api/core/auth/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/auth"
)

type AuthController struct {
	AuthService auth.AuthService
}

func (this AuthController) Login(request *requests.LoginRequest) (string, error) {
	// TODO: validate credentials
	return this.AuthService.GenerateJwt("Sixto"), nil
}

// TODO: Add middleware to verify requests
