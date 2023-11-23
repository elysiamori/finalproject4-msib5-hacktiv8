package admin

import (
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// seed admin
func SeedAdmin(db *gorm.DB) {
	admin := models.User{}

	if db.Where("email = ?", "admin@gmail.com").First(&admin).Error != nil {
		hashPass, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		admin = models.User{
			Fullname: "Admin",
			Email:    "admin@gmail.com",
			Password: string(hashPass),
			Role:     "admin",
		}

		db.Create(&admin)
	}
}
