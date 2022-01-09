package product

type Product struct {
	// Id          string  `bson:"id"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Price       float64 `bson:"price"`
	Quantity    int32   `bson:"quantity"`
}
