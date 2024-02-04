package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := // "host=localhost user=postgres password=database dbname=gofiber01 port=5432 sslmode=disable"
		fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			config.Host,
			config.User,
			config.Password,
			config.DBName,
			config.Port,
			config.SSLMode,
		)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	fmt.Print("Connecting to database successfully!")
	return db, nil
}
