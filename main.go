package main

import (
	"github.com/Jdsatashi/goFiber02/models"
	"github.com/Jdsatashi/goFiber02/storage"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(c *fiber.Ctx) error {
	book := Book{}
	err := c.BodyParser(&book)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Invalid request"})
		return err
	}
	err = r.DB.Create(&book).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Bad request",
		})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book created",
	})
	return nil
}

func (r *Repository) GetBooks(c *fiber.Ctx) error {
	var books = &[]models.Books{}
	var err = r.DB.Find(books).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Bad request",
		})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get books success",
		"data":    books,
	})

	return nil
}

func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	bookModel := &models.Books{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Id cannot empty.",
		})
		return nil
	}
	err := r.DB.Delete(bookModel, id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Cannot delete item.",
		})
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Delete item success.",
	})
	return nil
}

func (r *Repository) GetBook(c *fiber.Ctx) error {
	bookModel := &models.Books{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Id cannot empty.",
		})
		return nil
	}
	// err := r.DB.First(&bookModel, id)
	err := r.DB.Where("id = ?", id).First(&bookModel).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Cannot get item.",
		})
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get item success.",
		"data":    bookModel,
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_books/:id", r.DeleteBook)
	api.Get("/get_book/:id", r.GetBook)
	api.Get("/books", r.GetBooks)
}

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

	err2 := models.MigrateBooks(db)
	if err2 != nil {
		return
	}

	if err != nil {
		log.Fatal("Can not connect to database!")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting Fiber app:", err)
	}
}
