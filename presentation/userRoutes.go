package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/user/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/dependencyInjector"
)

type UserRouter struct{}

func (this UserRouter) SignUp(c echo.Context) error {
	request := new(requests.SignUpRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	controller := new(dependencyInjector.ContainerBuilder).GetUserController()
	errorCode := controller.SignUp(request)
	if errorCode != nil {
		return BuildInvalidResponse(c, errorCode)
	}
	return c.JSON(http.StatusOK, nil)
}
