package seeders

import (
	"github.com/Jdsatashi/goFiber02/models"
	"gorm.io/gorm"
)

func RolesSeeding(db *gorm.DB) {

}

func RoleAssignPermissions(roleName string, perms []string, db *gorm.DB) error {
	role := &models.Roles{}
	var permissions []*models.Permissions
	// Check if role existed
	err := db.Where("name = ?", roleName).FirstOrCreate(&role).Error
	// Handle error
	if err != nil {
		return err
	}
	// Assign permission to group permissions
	for _, perm := range perms {
		// Create temp permission to get permissions from name perms slice
		permission := &models.Permissions{}
		err := db.Where("name = ?", perm).First(permission).Error
		// Handle error
		if err != nil {
			return err
		}
		// Append valid permission to permissions slice
		permissions = append(permissions, permission)
	}
	// Assign permissions to role model
	role.Permission = permissions
	// Handling assign permissions to role on database
	err = db.Model(role).Updates(&role).Error
	// Handle error
	if err != nil {
		return err
	}
	return nil
}
