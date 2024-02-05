package controllers

import (
	"fmt"
	"net/http"

	"github.com/Jdsatashi/goFiber02/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BooksController struct {
	DB *gorm.DB
}

func NewBooksController(db *gorm.DB) *BooksController {
	return &BooksController{DB: db}
}

func (ctr *BooksController) CreateBook(c *fiber.Ctx) error {
	book := models.Books{}
	err := c.BodyParser(&book)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Invalid request"})
		return err
	}
	err = ctr.DB.Create(&book).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Bad request",
		})
		return err
	}

	c.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "Book created",
		"data":    book,
	})
	return nil
}

func (ctr *BooksController) GetBooks(c *fiber.Ctx) error {
	var books = &[]models.Books{}
	var err = ctr.DB.Find(books).Error
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

func (ctr *BooksController) DeleteBook(c *fiber.Ctx) error {
	bookModel := &models.Books{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Id cannot empty.",
		})
		return nil
	}
	err := ctr.DB.Delete(bookModel, id)
	fmt.Printf("Error: %v\n", err)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Cannot delete item.",
		})
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Delete item success.",
	})
	return nil
}

func (ctr *BooksController) GetBook(c *fiber.Ctx) error {
	bookModel := &models.Books{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Id cannot empty.",
		})
		return nil
	}
	// err := ctr.DB.First(&bookModel, id)
	err := ctr.DB.Where("id = ?", id).First(&bookModel).Error
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

func (ctr *BooksController) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Id cannot empty.",
		})
	}
	bookModel := &models.Books{}
	// err := ctr.DB.First(&bookModel, id)
	err2 := ctr.DB.Where("id = ?", id).First(&bookModel).Error
	if err2 != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Cannot get item.",
		})
	}
	book := models.Books{}
	err := c.BodyParser(&book)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Invalid request"})
	}
	err3 := ctr.DB.Model(bookModel).Updates(&book).Error
	if err3 != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Bad request",
		})
	}
	c.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "Book updated",
		"data":    book,
	})
	return nil
}
