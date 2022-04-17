package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/validations/customErrors"
)

func BuildInvalidResponse(c echo.Context, errorCode error) error {
	if errorCode.Error() == customErrors.INTERNAL_ERROR {
		return c.JSON(http.StatusInternalServerError, customErrors.CustomResponse{ErrorCode: errorCode.Error()})
	}
	return c.JSON(http.StatusBadRequest, customErrors.CustomResponse{ErrorCode: errorCode.Error()})
}
