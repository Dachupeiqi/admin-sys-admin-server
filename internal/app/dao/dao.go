package dao

import (
	"ginAdminServer/internal/app/config"
	"ginAdminServer/internal/app/dao/dict"
	"ginAdminServer/internal/app/dao/menu"
	"ginAdminServer/internal/app/dao/role"
	"ginAdminServer/internal/app/dao/user"
	"ginAdminServer/internal/app/dao/util"
	"strings"

	"github.com/google/wire"
	"gorm.io/gorm"
) // end

// RepoSet repo injection
var RepoSet = wire.NewSet(
	util.TransSet,
	menu.MenuActionResourceSet,
	menu.MenuActionSet,
	menu.MenuSet,
	role.RoleMenuSet,
	role.RoleSet,
	user.UserRoleSet,
	user.UserSet,
	dict.DictDataSet,
	dict.DictTypeSet,
) // end

// Define repo type alias
type (
	TransRepo              = util.Trans
	MenuActionResourceRepo = menu.MenuActionResourceRepo
	MenuActionRepo         = menu.MenuActionRepo
	MenuRepo               = menu.MenuRepo
	RoleMenuRepo           = role.RoleMenuRepo
	RoleRepo               = role.RoleRepo
	UserRoleRepo           = user.UserRoleRepo
	UserRepo               = user.UserRepo
	DictDataRepo           = dict.DictDataRepo
	DictTypeRepo           = dict.DictTypeRepo
) // end

// Auto migration for given models
func AutoMigrate(db *gorm.DB) error {
	if dbType := config.C.Gorm.DBType; strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	return db.AutoMigrate(
		new(menu.MenuActionResource),
		new(menu.MenuAction),
		new(menu.Menu),
		new(role.RoleMenu),
		new(role.Role),
		new(user.UserRole),
		new(user.User),
		new(dict.DictData),
		new(dict.DictType),
	) // end
}
