package dto

type ProductInput struct {
	Title      string `json:"title" binding:"required"`
	Stock      int    `json:"stock" binding:"required,min=5"`
	Price      int    `json:"price" binding:"required,min=0,max=5000000"`
	CategoryID uint   `json:"category_id" binding:"required"`
}
