package product

type Product struct {
	Id          string  `bson:"_id" json:"id"`
	Name        string  `bson:"name" json:"name"`
	Description string  `bson:"description" json:"description"`
	Price       float64 `bson:"price" json:"price"`
	Quantity    int32   `bson:"quantity" json:"quantity"`
}
