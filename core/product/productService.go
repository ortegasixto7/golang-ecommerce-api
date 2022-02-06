package product

type IProductService interface {
	Save(Product)
	Update(Product)
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

func (this ProductService) Save(product Product) {
	this.Repository.Save(product)
}

func (this ProductService) Update(product Product) {
	this.Repository.Update(product)
}

func (this ProductService) Delete(id string) {
	this.Repository.Delete(id)
}

func (this ProductService) GetAll() []Product {
	return this.Repository.GetAll()
}

func (this ProductService) GetById(id string) Product {
	return this.Repository.GetById(id)
}
