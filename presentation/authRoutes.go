package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/auth/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/dependencyInjector"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/validations/customErrors"
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
		return this.buildInvalidResponse(c, error)
	}
	return c.JSON(http.StatusOK, echo.Map{"token": jwt})
}

func (this AuthRouter) buildInvalidResponse(c echo.Context, errorCode error) error {
	if errorCode.Error() == customErrors.INTERNAL_ERROR {
		return c.JSON(http.StatusInternalServerError, customErrors.CustomResponse{ErrorCode: errorCode.Error()})
	}
	return c.JSON(http.StatusBadRequest, customErrors.CustomResponse{ErrorCode: errorCode.Error()})
}
