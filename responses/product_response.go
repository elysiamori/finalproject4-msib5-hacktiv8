package responses

type ProductResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID uint   `json:"category_Id"`
	CreatedAt  string `json:"created_at"`
}

type ProductResponseUpdate struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID uint   `json:"Category_Id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type ProductResponseID struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Price      string `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID uint   `json:"Category_Id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type UpdateProductID struct {
	Product ProductResponseID `json:"product"`
}

type GetProductResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Products struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID uint   `json:"category_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
