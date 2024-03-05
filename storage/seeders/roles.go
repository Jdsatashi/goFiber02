package seeders

import (
	"fmt"
	"github.com/Jdsatashi/goFiber02/models"
	"gorm.io/gorm"
)

func RolesSeeding(db *gorm.DB) {

}

func RoleAssignPermissions(roleName string, perms []string, db *gorm.DB) error {
	role := &models.Roles{}
	var permissions []*models.Permissions
	// Check if role existed
	err := db.Where("name = ?", roleName).First(&role)
	// If not existed, create new one
	if err != nil {
		err2 := db.Create(&models.Roles{Name: roleName})
		fmt.Printf("Role not exist, create role: %v", roleName)
		// Handle log if got error
		if err2 != nil {
			return err2.Error
		}
	}
	db.Where("name = ?", roleName).First(&role)
	// Assign permissions to role
	for _, perm := range perms {
		permission := &models.Permissions{}
		err := db.Where("name = ?", perm).First(permission)
		if err != nil {
			return err.Error
		}
		permissions = append(permissions, permission)
	}
	role.Permission = permissions
	err = db.Model(role).Updates(&role)
	if err != nil {
		return err.Error
	}
	return nil
}
