package seeders

import (
	"github.com/Jdsatashi/GoFiber02/models"
	"github.com/Jdsatashi/goFiber02/models"
)

func PermissionsSeeding() {
	perms := &[]models.Permissions{}
}

func data(){
	perms0 := &models.Permissions{
		Name: "list_books"
		Description: "Allow user to view all books."
	}
	perms1 := &models.Permissions{
		Name: "create_books"
		Description: "Allow user create books."
	}
	perms2 := &models.Permissions{
		Name: "view_books"
		Description: "Allow user view detail books."
	}
	perms3 := &models.Permissions{
		Name: "edit_books"
		Description: "Allow user edit books."
	}
	perms4 := &models.Permissions{
		Name: "delete_books"
		Description: "Allow user delete books."
	}

}