package models

import (
	"time"
)

type Product struct {
	ID          uint                 `gorm:"primarykey" json:"id"`
	Title       string               `gorm:"not null" json:"title" form:"title" valid:"required-Your title is required"`
	Price       int                  `gorm:"not null" json:"price" form:"price" valid:"required-Your price is required"`
	Stock       int                  `gorm:"not null" json:"stock" form:"stock" valid:"required-Your stock is required"`
	CreatedAt   time.Time            `gorm:"not null;autoCreateTime" json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt   time.Time            `gorm:"not null;autoUpdateTime" json:"updated_at" time_format:"2006-01-02 15:04:05"`
	CategoryID  uint                 `json:"-"`
	Transaction []TransactionHistory `gorm:"foreignKey:ProductID"`
}
