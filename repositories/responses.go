package repositories

type UserResponse struct {
	UserCode string `json:"user_code"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type BookResponse struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Author    string `json:"author"`
	Title     string `gorm:"unique" json:"title"`
	Publisher string `json:"publisher"`
	WriterID  string `json:"writer_name"`
}
