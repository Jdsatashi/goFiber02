package models

import "gorm.io/gorm"

type Users struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Books struct {
	ID        uint    `gorm:"primary key; autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	var err = db.AutoMigrate(&Books{})
	return err
}

func MigrateUsers(db *gorm.DB) error {
	var err = db.AutoMigrate(&Users{})
	return err
}

type LoginResponse struct {
	Token string `json:"token"`
}
