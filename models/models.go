package models

import (
	"time"

	"gorm.io/gorm"
)

type TimeStamped struct {
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time
}

// Users User model for login infomation
type Users struct {
	UserCode    string         `gorm:"primaryKey;autoIncrement:false;type:varchar(10);unique" json:"usercode"`
	Email       string         `gorm:"type:varchar(50);unique;not null" json:"email"`
	Username    string         `gorm:"type:varchar(50);unique;not null" json:"username"`
	PhoneNumber string         `gorm:"type:varchar(18);unique;not null" json:"phone_number"`
	Password    string         `gorm:"not null" json:"password"`
	Permission  []*Permissions `gorm:"many2many:users_permissions"`
	Role        []*Roles       `gorm:"many2many:users_roles"`
	TimeStamped
}

// Permissions Create permissions for users
type Permissions struct {
	ID          int      `gorm:"primaryKey;unique;autoIncrement:true" json:"id"`
	Name        string   `gorm:"unique" json:"name"`
	Description string   `json:"description"`
	User        []*Users `gorm:"many2many:users_permissions"`
	Role        []*Roles `gorm:"many2many:roles_permissions"`
	TimeStamped
}

// Roles as permissions group
type Roles struct {
	ID          int            `gorm:"primaryKey;unique" json:"id"`
	Name        string         `gorm:"unique" json:"role_name"`
	Description string         `json:"description"`
	Permission  []*Permissions `gorm:"many2many:users_permissions"`
	User        []*Users       `gorm:"many2many:users_roles"`
}

// Books Book for testing data
type Books struct {
	ID        int    `gorm:"primaryKey;unique" json:"id"`
	Author    string `json:"author"`
	Title     string `gorm:"unique" json:"title"`
	Publisher string `json:"publisher"`
	WriterID  string `json:"writer_name"`
	Writer    Users  `gorm:"foreignKey:WriterID;references:Username"`
	TimeStamped
}

// LoginTokens Specific and add token to database
type LoginTokens struct {
	ID           int       `gorm:"primaryKey;unique" json:"id"`
	Token        string    `gorm:"unique" json:"token"`
	ClientIP     string    `gorm:"unique" json:"client_ip"`
	DeviceType   string    `json:"device_type"`
	NumberDevice int       `json:"number_of_devices"`
	UserCode     string    `json:"user_code"`
	User         Users     `gorm:"foreignKey:UserCode;references:UserCode"`
	TimeExpired  time.Time `json:"time_expired"`
	CreatedAt    time.Time `json:"created_at"`
}

// MigrateBooks Migrating to create table
func MigrateBooks(db *gorm.DB) error {
	var err = db.AutoMigrate(&Books{})
	return err
}

func MigrateUsers(db *gorm.DB) error {
	var err = db.AutoMigrate(&Users{})
	return err
}
