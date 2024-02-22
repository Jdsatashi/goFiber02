package models

import (
	"gorm.io/gorm"
	"time"
)

type TimeStamped struct {
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time
}

type Users struct {
	UserCode   string         `gorm:"primaryKey;autoIncrement:false;type:varchar(10)" json:"usercode"`
	Email      string         `gorm:"type:varchar(50);unique;not null" json:"email"`
	Username   string         `gorm:"type:varchar(50);unique;not null" json:"username"`
	Password   string         `json:"password"`
	Permission []*Permissions `gorm:"many2many:users_permissions"`
}

type Permissions struct {
	ID          int      `gorm:"primaryKey;" json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	User        []*Users `gorm:"many2many:users_permissions"`
}

type Books struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Author    string `json:"author"`
	Title     string `gorm:"unique" json:"title"`
	Publisher string `json:"publisher"`
	WriterID  string `json:"writer_name"`
	Writer    Users  `gorm:"foreignKey:WriterID;references:Username"`
}

func MigrateBooks(db *gorm.DB) error {
	var err = db.AutoMigrate(&Books{})
	return err
}

func MigrateUsers(db *gorm.DB) error {
	var err = db.AutoMigrate(&Users{})
	return err
}
