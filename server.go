package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requestModels"
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

		return c.JSON(http.StatusOK, request)
	})

	server.Logger.Fatal(server.Start(":1323"))
}
