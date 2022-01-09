package product

import "fmt"

type IProductService interface {
	Save(*Product)
	Update(*Product)
	Delete(id string)
	GetAll() []Product
	GetById(id string) Product
}

type ProductService struct {
	Repository IProductPersistence
}

func NewProductService(repository IProductPersistence) *ProductService {
	return &ProductService{Repository: repository}
}

func (service ProductService) Save(product Product) {
	service.Repository.Save(product)
	fmt.Print("Product Saved")
}

func (p ProductService) Update(product Product) {
	fmt.Print("Product Updated")
}

func (p ProductService) Delete(id string) {
	fmt.Print("Product Deleted")
}

func (p ProductService) GetAll() []Product {
	return []Product{}
}

func (p ProductService) GetById(id string) Product {
	return Product{}
}
