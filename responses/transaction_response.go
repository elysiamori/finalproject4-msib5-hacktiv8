package responses

type TransactionBill struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type CTResponse struct {
	Message         string          `json:"message"`
	TransactionBill TransactionBill `json:"transaction_bill"`
}

type TransactionHistory struct {
	ID         uint     `gorm:"primarykey" json:"id"`
	ProductID  uint     `gorm:"not null" json:"product_id"`
	UserID     uint     `gorm:"not null" json:"user_id"`
	Quantity   int      `gorm:"not null" json:"quantity"`
	TotalPrice int      `gorm:"not null" json:"total_price"`
	Product    Products `json:"Product"`
}

type TransactionHistories struct {
	ID         uint     `gorm:"primarykey" json:"id"`
	ProductID  uint     `gorm:"not null" json:"product_id"`
	UserID     uint     `gorm:"not null" json:"user_id"`
	Quantity   int      `gorm:"not null" json:"quantity"`
	TotalPrice int      `gorm:"not null" json:"total_price"`
	Product    Products `json:"Product"`
	User       Users    `json:"User"`
}
