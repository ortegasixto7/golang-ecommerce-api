package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/dependencyInjector"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/validations/customErrors"
)

type AdminRouter struct{}

func (this AdminRouter) CreateAdminUser(c echo.Context) error {
	controller := new(dependencyInjector.ContainerBuilder).GetAdminController()
	errorCode := controller.CreateAdminUser()
	if errorCode != nil {
		return this.buildInvalidResponse(c, errorCode)
	}
	return c.JSON(http.StatusOK, nil)
}

func (this AdminRouter) buildInvalidResponse(c echo.Context, errorCode error) error {
	if errorCode.Error() == customErrors.INTERNAL_ERROR {
		return c.JSON(http.StatusInternalServerError, customErrors.CustomResponse{ErrorCode: errorCode.Error()})
	}
	return c.JSON(http.StatusBadRequest, customErrors.CustomResponse{ErrorCode: errorCode.Error()})
}
