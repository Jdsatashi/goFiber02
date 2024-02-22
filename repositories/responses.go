package repositories

type TokenResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	UserCode string `json:"user_code"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
