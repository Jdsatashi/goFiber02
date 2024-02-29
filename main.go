package main

import (
	"fmt"
	"github.com/Jdsatashi/goFiber02/storage/seeders"
	"log"
	"os"

	"github.com/Jdsatashi/goFiber02/models"
	"github.com/Jdsatashi/goFiber02/routes"
	"github.com/Jdsatashi/goFiber02/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Can not connect to database!")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("Can not migrate to database!")
	} else {
		fmt.Println("\nMigrating to Books")
	}
	err = models.MigrateUsers(db)
	if err != nil {
		log.Fatal("Can not migrate to database!")
	} else {
		fmt.Println("Migrating to Users")
	}
	seeders.PermissionsSeeding(db)
	app := fiber.New()
	routes.SetupRouting(app, db)
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting Fiber app:", err)
	}
}
