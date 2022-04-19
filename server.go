package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/dependencyInjector"
	"github.com/ortegasixto7/echo-go-supermarket-api/persistence/mongoDb"
	"github.com/ortegasixto7/echo-go-supermarket-api/presentation"
)

// Middlewares
func adminUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		controller := new(dependencyInjector.ContainerBuilder).GetAuthController()
		token := c.Request().Header.Get("Authorization")
		splitToken := strings.Split(token, "Bearer ")
		token = splitToken[1]
		userId := controller.AuthService.DecodeJwt(token)
		fmt.Println(userId)
		if userId == "" {
			return echo.ErrUnauthorized
		}

		adminUser, isEmpty := controller.AdminPersistence.GetById(userId)
		if isEmpty {
			return echo.ErrUnauthorized
		}

		if adminUser.Role != "ADMIN" {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func main() {
	loadEnvFileError := godotenv.Load(".env")
	if loadEnvFileError != nil {
		log.Fatalf("Error loading .env file")
	}
	server := echo.New()
	mongoDb.Setup()
	server.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "API is working!")
	})
	server.PUT("/products/add-quantity", presentation.ProductRouter{}.AddQuantity, adminUserMiddleware)
	server.PUT("/products", presentation.ProductRouter{}.Update, adminUserMiddleware)
	server.POST("/products", presentation.ProductRouter{}.Create, adminUserMiddleware)
	server.GET("/products", presentation.ProductRouter{}.GetAll)
	server.GET("/products/:id", presentation.ProductRouter{}.GetById)
	server.POST("/admin/create-admin-user", presentation.AdminRouter{}.CreateAdminUser)
	server.POST("/auth/login", presentation.AuthRouter{}.Login)
	server.POST("/auth/sign-up", presentation.UserRouter{}.SignUp)
	server.Logger.Fatal(server.Start(":1323"))
}
