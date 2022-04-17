package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/dependencyInjector"
)

type AdminRouter struct{}

func (this AdminRouter) CreateAdminUser(c echo.Context) error {
	controller := new(dependencyInjector.ContainerBuilder).GetAdminController()
	errorCode := controller.CreateAdminUser()
	if errorCode != nil {
		return BuildInvalidResponse(c, errorCode)
	}
	return c.JSON(http.StatusOK, nil)
}
