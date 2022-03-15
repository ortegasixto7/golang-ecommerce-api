package admin

import (
	"golang.org/x/crypto/bcrypt"
)

type AdminController struct {
	AdminService AdminService
}

func (this AdminController) CreateAdminUser() (errorCode string) {

	var adminId = "62201eae6dbbd783dccb3c6b"

	adminUser := this.AdminService.GetById(adminId)
	if adminUser.Id == "" {

		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)

		adminUser.Id = adminId
		adminUser.FirstName = "Admin"
		adminUser.LastName = "User"
		adminUser.Password = string(encryptedPassword)
		adminUser.Role = "ADMIN"
		adminUser.Username = "admin"
		this.AdminService.Save(adminUser)
	}

	return ""
}
