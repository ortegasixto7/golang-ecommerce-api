package user

type IUserPersistence interface {
	Save(User)
	Update(User)
	GetById(id string) (User, bool)
	GetByUserName(username string) (User, bool)
}
