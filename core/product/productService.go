package product

type ProductService interface {
	Save(*Product)
	Update(*Product)
	Delete(id string)
	GetAll() []Product
	GetById(id string) Product
}
