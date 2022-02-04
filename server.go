package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requestModels"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requestValidations"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/validations/customErrors"
)

func main() {
	server := echo.New()
	server.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, World!")
	})

	server.POST("/products", func(c echo.Context) error {
		request := new(requestModels.CreateProductRequest)
		if err := c.Bind(request); err != nil {
			return err
		}

		requestError, errorCode := new(requestValidations.CreateProductRequestValidation).Validate(*request)
		if requestError != "" {
			return buildInvalidResponse(c, requestError, errorCode)
		}

		return c.JSON(http.StatusOK, request)
	})

	server.Logger.Fatal(server.Start(":1323"))
}

func buildInvalidResponse(c echo.Context, requestError customErrors.RequestErrorEnum, errorCode customErrors.ErrorCodeEnum) error {
	if requestError == customErrors.BAD_REQUEST {
		return c.JSON(http.StatusBadRequest, customErrors.CustomResponse{ErrorCode: string(errorCode)})
	}
	if requestError == customErrors.NOT_FOUND {
		return c.JSON(http.StatusNotFound, customErrors.CustomResponse{ErrorCode: string(errorCode)})
	}
	return c.JSON(http.StatusNotFound, customErrors.CustomResponse{ErrorCode: string(errorCode)})
}
