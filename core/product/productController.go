package product

import "github.com/ortegasixto7/echo-go-supermarket-api/core/product/requestModels"

type ProductController struct {
	ProductService ProductService
}

func (controller ProductController) Save(request requestModels.CreateProductRequest) {
	product := Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Quantity:    0}
	controller.ProductService.Save(product)
}
