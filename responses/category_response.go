package responses

type CategoryResponse struct {
	ID          uint                 `json:"id"`
	Type        string               `json:"type"`
	SoldProduct int                  `json:"sold_amount_product"`
	CreatedAt   string               `json:"created_at"`
	UpdatedAt   string               `json:"updated_at"`
	Products    []GetProductResponse `json:"Products"`
}

type CategoryResponseAdd struct {
	ID          uint   `json:"id"`
	Type        string `json:"type"`
	SoldProduct int    `json:"sold_amount_product"`
	CreatedAt   string `json:"created_at"`
}

type CategoryResponseUpdate struct {
	ID          uint   `json:"id"`
	Type        string `json:"type"`
	SoldProduct int    `json:"sold_amount_product"`
	UpdatedAt   string `json:"updated_at"`
}
