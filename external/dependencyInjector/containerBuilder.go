package dependencyInjector

import (
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product"
	"github.com/ortegasixto7/echo-go-supermarket-api/persistence/mongoDb"
)

type ContainerBuilder struct{}

func (container ContainerBuilder) getProductService() *product.ProductService {
	return &product.ProductService{Repository: &mongoDb.MongoProductPersistence{}}
}

func (container ContainerBuilder) GetProductController() *product.ProductController {
	return &product.ProductController{ProductService: *container.getProductService()}
}
