package admin

type IAdminPersistence interface {
	Save(AdminUser)
	Update(AdminUser)
	GetById(id string) (AdminUser, bool)
	GetByUserName(username string) (AdminUser, bool)
}
