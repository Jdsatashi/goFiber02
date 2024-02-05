package routes

import (
	"github.com/Jdsatashi/goFiber02/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRouting(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api/v1")
	controller := controllers.NewBooksController(db)

	api.Post("/create_books", controller.CreateBook)
	api.Delete("/delete_books/:id", controller.DeleteBook)
	api.Get("/get_book/:id", controller.GetBook)
	api.Put("/get_book/:id", controller.UpdateBook)
	api.Get("/books", controller.GetBooks)
}
