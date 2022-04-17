package admin

import (
	"golang.org/x/crypto/bcrypt"
)

type AdminController struct {
	AdminPersistence IAdminPersistence
}

func (this AdminController) CreateAdminUser() error {
	var adminId = "62201eae6dbbd783dccb3c6b"
	adminUser, isNil := this.AdminPersistence.GetById(adminId)
	if isNil {
		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
		adminUser.Id = adminId
		adminUser.FirstName = "Admin"
		adminUser.LastName = "User"
		adminUser.Password = string(encryptedPassword)
		adminUser.Role = "ADMIN"
		adminUser.Username = "admin"
		this.AdminPersistence.Save(adminUser)
	}
	return nil
}
