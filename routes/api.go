package routes

import (
	"github.com/Jdsatashi/goFiber02/config"
	"github.com/Jdsatashi/goFiber02/controllers"
	"github.com/Jdsatashi/goFiber02/middlewares"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRouting(app *fiber.App, db *gorm.DB) {
	jwt := middlewares.NewAuthMiddleware(config.Secret)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Home page Go Fiber 2",
		})
	})
	api := app.Group("/api/v1")
	bookCtrl := controllers.NewBooksController(db)
	authCtrl := controllers.NewAuthController(db)
	userCtrl := controllers.NewUsersController(db)
	// Book api routes
	api.Post("/create_books", bookCtrl.CreateBook)
	api.Delete("/delete_books/:id", bookCtrl.DeleteBook)
	api.Get("/get_book/:id", bookCtrl.GetBook)
	api.Put("/get_book/:id", bookCtrl.UpdateBook)
	api.Get("/books", bookCtrl.GetBooks)

	api.Post("/register", userCtrl.UserCreate)
	api.Post("/login", authCtrl.Login)
	api.Get("/middleware", jwt, controllers.Protected)

	api.Post("/account/create", jwt, userCtrl.UserCreate)
	api.Get("/accounts/", jwt, userCtrl.UserList)
	api.Get("/account/:user_code", jwt, userCtrl.UserDetail)
	api.Put("/account/:user_code/edit", jwt, userCtrl.UserUpdate)
	api.Delete("account/:user_code/delete", jwt, userCtrl.UserDelete)
}
