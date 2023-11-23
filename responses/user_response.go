package responses

type UserResponse struct {
	ID        uint   `json:"id"`
	Fullname  string `json:"full_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
}

type Users struct {
	ID        uint   `json:"id"`
	Fullname  string `json:"full_name"`
	Email     string `json:"email"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
