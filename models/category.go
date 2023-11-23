package models

import (
	"time"
)

type Category struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Type        string    `gorm:"not null" json:"type" form:"type"`
	SoldProduct int       `gorm:"default:0" json:"sold_amount_product"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime" json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt   time.Time `gorm:"not null;autoUpdateTime" json:"updated_at" time_format:"2006-01-02 15:04:05"`
	Products    []Product `gorm:"foreignKey:CategoryID" json:"Products"`
}
