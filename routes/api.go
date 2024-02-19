package routes

import (
	"github.com/Jdsatashi/goFiber02/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRouting(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api/v1")
	bookCtrl := controllers.NewBooksController(db)
	userCtrl := controllers.NewUsersController(db)

	api.Post("/create_books", bookCtrl.CreateBook)
	api.Delete("/delete_books/:id", bookCtrl.DeleteBook)
	api.Get("/get_book/:id", bookCtrl.GetBook)
	api.Put("/get_book/:id", bookCtrl.UpdateBook)
	api.Get("/books", bookCtrl.GetBooks)

	api.Post("/register", userCtrl.Register)
	api.Post("/login", userCtrl.Login)
	api.Get("/middleware", controllers.Protected)
}
