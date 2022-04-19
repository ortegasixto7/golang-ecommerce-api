package admin

import (
	"github.com/ortegasixto7/echo-go-supermarket-api/core/shared"
	"golang.org/x/crypto/bcrypt"
)

type AdminController struct {
	AdminPersistence IAdminPersistence
}

func (this AdminController) CreateAdminUser() error {
	var adminId = "62201eae6dbbd783dccb3c6b"
	adminUser, isEmpty := this.AdminPersistence.GetById(adminId)
	if isEmpty {
		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
		adminUser.Id = adminId
		adminUser.FirstName = "Admin"
		adminUser.LastName = "User"
		adminUser.Password = string(encryptedPassword)
		adminUser.Role = shared.ADMIN
		adminUser.Username = "admin"
		this.AdminPersistence.Save(adminUser)
	}
	return nil
}
