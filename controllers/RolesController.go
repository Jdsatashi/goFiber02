package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RolesController struct {
	DB *gorm.DB
}

func NewRolesController(db *gorm.DB) *RolesController {
	return &RolesController{DB: db}
}

func (ctr *RolesController) RoleCreate(c *fiber.Ctx) error {
	return nil
}
