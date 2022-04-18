package dependencyInjector

import (
	"github.com/ortegasixto7/echo-go-supermarket-api/core/admin"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/auth"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product"
	externalAuth "github.com/ortegasixto7/echo-go-supermarket-api/external/auth"
	"github.com/ortegasixto7/echo-go-supermarket-api/persistence/mongoDb"
)

type ContainerBuilder struct{}

func (container ContainerBuilder) getAdminPersistence() admin.IAdminPersistence {
	return mongoDb.MongoAdminPersistence{}
}

func (container ContainerBuilder) getProductPersistence() product.IProductPersistence {
	return mongoDb.MongoProductPersistence{}
}

func (container ContainerBuilder) getAuthService() *externalAuth.AuthService {
	return &externalAuth.AuthService{}
}

func (container ContainerBuilder) GetProductController() *product.ProductController {
	return &product.ProductController{ProductPersistence: container.getProductPersistence()}
}

func (container ContainerBuilder) GetAdminController() *admin.AdminController {
	return &admin.AdminController{AdminPersistence: container.getAdminPersistence()}
}

func (container ContainerBuilder) GetAuthController() *auth.AuthController {
	return &auth.AuthController{AuthService: *container.getAuthService(), AdminPersistence: container.getAdminPersistence()}
}
