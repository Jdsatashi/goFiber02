package controllers

import (
	"github.com/Jdsatashi/goFiber02/handlers"
	"github.com/Jdsatashi/goFiber02/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

var userHandler = handlers.UserHandler{}

type UsersController struct {
	DB *gorm.DB
}

func NewUsersController(db *gorm.DB) *UsersController {
	return &UsersController{DB: db}
}

func (ctr *UsersController) UserCreate(c *fiber.Ctx) error {
	user := &models.Users{}
	err := c.BodyParser(user)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message": "Invalid request",
				"error":   err.Error(),
			})
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Cannot encrypting password",
		})
		return err
	}
	user.Password = string(hashedPassword)
	err = ctr.DB.Create(&user).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Bad request",
		})
		return err
	}
	response := userHandler.ToUserResponse(*user)
	c.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "User created",
		"data":    response,
	})
	return nil
}

func (ctr *UsersController) UserList(c *fiber.Ctx) error {
	users := &[]models.Users{}
	err := ctr.DB.Find(users).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Bad request",
			"error":   err.Error(),
		})
		return err
	}
	responses := userHandler.ToUsersResponse(*users)
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Users retrieved",
		"data":    responses,
	})
	return nil
}

func (ctr *UsersController) UserDetail(c *fiber.Ctx) error {
	user := &models.Users{}
	userId := c.Params("user_code")
	if userId == "" {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "User code cannot empty.",
		})
	}
	err := ctr.DB.Where("user_code = ?", userId).First(&user).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Cannot find item with user code = " + userId,
		})
	}
	response := userHandler.ToUserResponse(*user)
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get successfully user: " + user.Username,
		"data":    response,
	})
	return nil
}

func (ctr *UsersController) UserUpdate(c *fiber.Ctx) error {
	userModel := &models.Users{}
	userId := c.Params("user_code")
	if userId == "" {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "User code cannot empty.",
		})
	}
	err := ctr.DB.Where("user_code = ?", userId).First(&userModel).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Cannot find item with user code = " + userId,
			"error":   err.Error(),
		})
	}
	updateUser := &models.Users{}
	err = c.BodyParser(updateUser)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message": "Invalid request",
				"error":   err.Error(),
			})
		return err
	}
	err = ctr.DB.Model(&userModel).Updates(&updateUser).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Bad request",
			"error":   err.Error(),
		})
	}
	response := userHandler.ToUserResponse(*updateUser)
	c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Update successfully user " + updateUser.Username,
		"data":    response,
	})
	return nil
}

func (ctr *UsersController) UserDelete(c *fiber.Ctx) error {
	userModel := &models.Users{}
	userId := c.Params("user_code")
	if userId == "" {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "User code cannot empty.",
		})
	}
	err := ctr.DB.Where("user_code = ?", userId).Delete(&userModel).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Cannot find item with user code = " + userId,
			"error":   err.Error(),
		})
	}
	c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Delete user successfully",
	})
	return nil
}
