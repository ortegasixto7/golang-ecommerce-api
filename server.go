package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/dependencyInjector"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/validations/customErrors"
)

func main() {
	server := echo.New()
	server.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, World!")
	})

	server.POST("/products", func(c echo.Context) error {
		request := new(requests.CreateProductRequest)
		if err := c.Bind(request); err != nil {
			return err
		}
		productController := new(dependencyInjector.ContainerBuilder).GetProductController()
		requestError, errorCode := productController.Create(request)
		if requestError != "" {
			return buildInvalidResponse(c, requestError, errorCode)
		}
		return c.JSON(http.StatusOK, request)
	})

	server.Logger.Fatal(server.Start(":1323"))
}

func buildInvalidResponse(c echo.Context, requestError string, errorCode string) error {
	if requestError == string(customErrors.BAD_REQUEST) {
		return c.JSON(http.StatusBadRequest, customErrors.CustomResponse{ErrorCode: string(errorCode)})
	}
	if requestError == string(customErrors.NOT_FOUND) {
		return c.JSON(http.StatusNotFound, customErrors.CustomResponse{ErrorCode: string(errorCode)})
	}
	return c.JSON(http.StatusNotFound, customErrors.CustomResponse{ErrorCode: string(errorCode)})
}
