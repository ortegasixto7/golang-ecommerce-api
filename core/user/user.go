package user

type User struct {
	Id        string `bson:"_id" json:"id"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Username  string `bson:"username" json:"username"`
	Password  string `bson:"password" json:"password"`
	Role      string `bson:"role" json:"role"`
}
