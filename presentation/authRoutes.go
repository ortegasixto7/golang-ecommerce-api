package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/auth/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/dependencyInjector"
)

type AuthRouter struct{}

func (this AuthRouter) Login(c echo.Context) error {
	request := new(requests.LoginRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	controller := new(dependencyInjector.ContainerBuilder).GetAuthController()
	jwt, error := controller.Login(request)
	if error != nil {
		return BuildInvalidResponse(c, error)
	}
	return c.JSON(http.StatusOK, echo.Map{"token": jwt})
}
