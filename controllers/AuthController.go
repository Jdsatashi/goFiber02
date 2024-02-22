package controllers

import (
	"github.com/Jdsatashi/goFiber02/config"
	"github.com/Jdsatashi/goFiber02/models"
	repo "github.com/Jdsatashi/goFiber02/repositories"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ctr *AuthController) Login(c *fiber.Ctx) error {
	loginRequest := &repo.LoginRequest{}
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
	date := time.Hour * 12
	claims := jtoken.MapClaims{
		"UserCode": user.UserCode,
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
			"UserCode": user.UserCode,
			"email":    user.Email,
			"username": user.Username,
		},
	})
}

func Protected(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)
	username := claims["username"].(string)
	return c.SendString("Welcome " + email + " " + username)
}
