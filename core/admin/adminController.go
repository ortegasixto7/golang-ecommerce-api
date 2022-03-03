package admin

type AdminController struct {
	AdminService AdminService
}

func (this AdminController) CreateAdminUser() (errorCode string) {

	var adminId = "62201eae6dbbd783dccb3c6b"

	adminUser := this.AdminService.GetById(adminId)
	if adminUser.Id == "" {
		adminUser.Id = adminId
		adminUser.FirstName = "Admin"
		adminUser.LastName = "User"
		// TODO: Encrypt password here
		adminUser.Password = "123456"
		adminUser.Role = "ADMIN"
		adminUser.Username = "admin"
		this.AdminService.Save(adminUser)
	}

	return ""
}
