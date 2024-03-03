package seeders

import (
	"errors"
	"fmt"
	"github.com/Jdsatashi/goFiber02/models"
	"gorm.io/gorm"
	"log"
)

func PermissionsSeeding(db *gorm.DB) {
	_ = AutoCreateCRUD("books", db)
	_ = AutoCreateCRUD("users", db)
	_ = CreatePerm([]string{"own_user_view", "Allow user to view own account"}, db)
	_ = CreatePerm([]string{"own_user_edit", "Allow user to edit own account"}, db)
}

// CreatePerm Create specific permission
func CreatePerm(perm []string, db *gorm.DB) error {
	newPerm := &models.Permissions{
		Name:        perm[0],
		Description: perm[1],
	}
	if newPerm.Name == "" {
		return errors.New("permission's name can not be empty")
	}
	fmt.Printf("Permissions to add: %v\n", perm)
	err := db.Where("name = ?", newPerm.Name).FirstOrCreate(&newPerm).Error
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		log.Fatalf("Failed when auto-create %v permission!", newPerm.Name)
	}
	return nil
}

// AutoCreateCRUD will auto create CRUD permissions for a function
func AutoCreateCRUD(perm string, db *gorm.DB) error {
	if perm == "" {
		return errors.New("permission's group name can not be empty")
	}
	var groupPerms []models.Permissions
	perms0 := &models.Permissions{
		Name:        "list_" + perm,
		Description: "Allow user to view all " + perm,
	}
	groupPerms = append(groupPerms, *perms0)
	perms1 := &models.Permissions{
		Name:        "create_" + perm,
		Description: "Allow user create " + perm,
	}
	groupPerms = append(groupPerms, *perms1)
	perms2 := &models.Permissions{
		Name:        "view_" + perm,
		Description: "Allow user view detail " + perm,
	}
	groupPerms = append(groupPerms, *perms2)
	perms3 := &models.Permissions{
		Name:        "edit_" + perm,
		Description: "Allow user edit " + perm,
	}
	groupPerms = append(groupPerms, *perms3)
	perms4 := &models.Permissions{
		Name:        "delete_" + perm,
		Description: "Allow user delete " + perm,
	}
	groupPerms = append(groupPerms, *perms4)
	for _, perm := range groupPerms {
		fmt.Printf("Permissions to add: %v\n", perm)
		err := db.Where("name = ?", perm.Name).FirstOrCreate(&perm).Error
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			log.Fatalf("Failed when auto-create %v permission!", perm.Name)
		}
	}
	return nil
}
