package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uint                 `gorm:"primaryKey" json:"id"`
	Fullname    string               `gorm:"not null" json:"full_name" form:"full_name" valid:"required-Your full name is required"`
	Email       string               `gorm:"not null uniqueIndex" json:"email" form:"email" valid:"email,required-Please enter a valid email address"`
	Password    string               `gorm:"not null" json:"password" form:"password" valid:"required-Password is required, minstringlength(6)-Password should be more than 6 chars"`
	Balance     int                  `gorm:"default:0" json:"balance" valid:"required,min=0,max=100000000"`
	Role        string               `gorm:"default:customer" json:"role" valid:"required"`
	CreatedAt   time.Time            `gorm:"not null;autoCreateTime" json:"created_at"`
	Transaction []TransactionHistory `gorm:"foreignKey:UserID"`
}

// verify password
func VerifyPassword(password, hashPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))

}
