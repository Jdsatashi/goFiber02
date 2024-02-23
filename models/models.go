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
	UserCode   string         `gorm:"primaryKey;autoIncrement:false;type:varchar(10);unique" json:"usercode"`
	Email      string         `gorm:"type:varchar(50);unique;not null" json:"email"`
	Username   string         `gorm:"type:varchar(50);unique;not null" json:"username"`
	Password   string         `json:"password"`
	Permission []*Permissions `gorm:"many2many:users_permissions"`
	TimeStamped
}

type Permissions struct {
	ID          int      `gorm:"primaryKey;unique" json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	User        []*Users `gorm:"many2many:users_permissions"`
	TimeStamped
}

type Books struct {
	ID        int    `gorm:"primaryKey;unique" json:"id"`
	Author    string `json:"author"`
	Title     string `gorm:"unique" json:"title"`
	Publisher string `json:"publisher"`
	WriterID  string `json:"writer_name"`
	Writer    Users  `gorm:"foreignKey:WriterID;references:Username"`
	TimeStamped
}

type LoginToken struct {
	ID           int       `gorm:"primaryKey;unique" json:"id"`
	Token        string    `gorm:"unique" json:"token"`
	ClientIP     string    `gorm:"unique" json:"client_ip"`
	NumberDevice int       `json:"number_of_devices"`
	UserCode     string    `json:"user_code"`
	User         Users     `gorm:"foreignKey:UserCode;references:UserCode"`
	TimeExpired  time.Time `json:"time_expired"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

func MigrateBooks(db *gorm.DB) error {
	var err = db.AutoMigrate(&Books{})
	return err
}

func MigrateUsers(db *gorm.DB) error {
	var err = db.AutoMigrate(&Users{})
	return err
}
