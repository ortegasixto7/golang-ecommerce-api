package dependencyInjector

import (
	"github.com/ortegasixto7/echo-go-supermarket-api/core/admin"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/auth"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product"
	externalAuth "github.com/ortegasixto7/echo-go-supermarket-api/external/auth"
	"github.com/ortegasixto7/echo-go-supermarket-api/persistence/mongoDb"
)

type ContainerBuilder struct{}

func (container ContainerBuilder) getAdminService() *admin.AdminService {
	return &admin.AdminService{Persistence: &mongoDb.MongoAdminPersistence{}}
}

func (container ContainerBuilder) getProductService() *product.ProductService {
	return &product.ProductService{Persistence: &mongoDb.MongoProductPersistence{}}
}

func (container ContainerBuilder) getAuthService() *externalAuth.AuthService {
	return &externalAuth.AuthService{}
}

func (container ContainerBuilder) GetProductController() *product.ProductController {
	return &product.ProductController{ProductService: *container.getProductService()}
}

func (container ContainerBuilder) GetAdminController() *admin.AdminController {
	return &admin.AdminController{AdminService: *container.getAdminService()}
}

func (container ContainerBuilder) GetAuthController() *auth.AuthController {
	return &auth.AuthController{AuthService: *container.getAuthService(), AdminService: *container.getAdminService()}
}
