package models

type TransactionHistory struct {
	ID         uint    `gorm:"primarykey" json:"id"`
	UserID     uint    `gorm:"not null" json:"user_id"`
	ProductID  uint    `gorm:"not null" json:"product_id"`
	Quantity   int     `gorm:"not null" json:"quantity" valid:"required"`
	TotalPrice int     `gorm:"not null" json:"total" valid:"required"`
	User       User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Product    Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
