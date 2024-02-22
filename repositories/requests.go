package repositories

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PasswordUpdate struct {
	Password string `json:"password"`
}

type BookCreate struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	Writer    string `json:"writer"`
}
