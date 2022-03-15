package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/persistence/mongoDb"
	"github.com/ortegasixto7/echo-go-supermarket-api/presentation"
)

func main() {
	server := echo.New()
	mongoDb.Setup()
	server.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "API is working!")
	})
	server.PUT("/products/add-quantity", func(context echo.Context) error {
		return presentation.ProductRouter{}.AddQuantity(context)
	})
	server.PUT("/products", func(context echo.Context) error {
		return presentation.ProductRouter{}.Update(context)
	})
	server.POST("/products", func(context echo.Context) error {
		return presentation.ProductRouter{}.Create(context)
	})
	server.GET("/products", func(context echo.Context) error {
		return presentation.ProductRouter{}.GetAll(context)
	})
	server.GET("/products/:id", func(context echo.Context) error {
		return presentation.ProductRouter{}.GetById(context)
	})
	server.POST("/admin/create-admin-user", func(context echo.Context) error {
		return presentation.AdminRouter{}.CreateAdminUser(context)
	})
	server.POST("/auth/login", func(context echo.Context) error {
		return presentation.AuthRouter{}.Login(context)
	})
	server.Logger.Fatal(server.Start(":1323"))
}
