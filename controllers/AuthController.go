package controllers

import (
	"fmt"
	"github.com/Jdsatashi/goFiber02/config"
	"github.com/Jdsatashi/goFiber02/models"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type UsersController struct {
	DB *gorm.DB
}

func NewUsersController(db *gorm.DB) *UsersController {
	return &UsersController{DB: db}
}

func (ctr *UsersController) Register(c *fiber.Ctx) error {
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
	c.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "User created",
		"data": &fiber.Map{
			"username": user.Username,
			"email":    user.Email,
		},
	})
	return nil
}

func (ctr *UsersController) Login(c *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)
	user := &models.Users{}
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err := ctr.DB.Where("email = ?", loginRequest.Email).First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Password incorrect",
		})
		return err
	}
	date := time.Hour * 24
	claims := jtoken.MapClaims{
		"ID":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(date * 1).Unix(),
	}
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"token": t,
		"user": fiber.Map{
			"ID":       user.ID,
			"email":    user.Email,
			"username": user.Username,
		},
	})
}

func Protected(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token)
	fmt.Printf("\nUser is %v", user)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)
	username := claims["username"].(string)
	return c.SendString("Welcome " + email + " " + username)
}
