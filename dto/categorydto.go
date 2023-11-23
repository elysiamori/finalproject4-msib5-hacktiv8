package dto

type CategoryInput struct {
	Type string `json:"type" binding:"required"`
}
