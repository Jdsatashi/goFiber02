package seeders

import (
	"fmt"
	"github.com/Jdsatashi/goFiber02/models"
	"gorm.io/gorm"
	"log"
)

func PermissionsSeeding(db *gorm.DB) {
	AutoCreate("books", db)
	AutoCreate("users", db)
}

func AutoCreate(perm string, db *gorm.DB) {
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
		err := db.Where(models.Permissions{Name: perm.Name}).FirstOrCreate(&perm).Error
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			log.Fatalf("Failed when auto-create %v permission!", perm.Name)
		}
	}
}
