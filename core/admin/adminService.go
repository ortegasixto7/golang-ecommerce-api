package admin

type IAdminService interface {
	Save(AdminUser)
	Update(AdminUser)
	GetById(id string) AdminUser
}

type AdminService struct {
	Persistence IAdminPersistence
}

func NewAdminService(persistence IAdminPersistence) *AdminService {
	return &AdminService{Persistence: persistence}
}

func (this AdminService) Update(data AdminUser) {
	this.Persistence.Update(data)
}

func (this AdminService) Save(data AdminUser) {
	this.Persistence.Save(data)
}

func (this AdminService) GetById(id string) AdminUser {
	return this.Persistence.GetById(id)
}