package remoterpc

import (
	"ginAdminServer/internal/app/config"
	"ginAdminServer/internal/app/dao"
	"ginAdminServer/internal/app/dao/dict"
	"ginAdminServer/internal/app/dao/menu"
	"ginAdminServer/internal/app/dao/role"
	"ginAdminServer/internal/app/dao/user"
	"ginAdminServer/internal/app/dao/util"
	"ginAdminServer/internal/app/service"
	"ginAdminServer/pkg/errors"
	"ginAdminServer/pkg/gormx"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

import (
	_ "ginAdminServer/internal/app/swagger"
)

// Injectors from wire.go:

func BuildInjectorForTalent() (*InjectorForTalent, func(), error) {

	db, cleanup2, err := InitGormDB()
	if err != nil {
		cleanup2()
		return nil, nil, err
	}
	roleRepo := &role.RoleRepo{
		DB: db,
	}
	roleMenuRepo := &role.RoleMenuRepo{
		DB: db,
	}
	menuActionResourceRepo := &menu.MenuActionResourceRepo{
		DB: db,
	}
	userRepo := &user.UserRepo{
		DB: db,
	}
	userRoleRepo := &user.UserRoleRepo{
		DB: db,
	}
	dictDataRepo := &dict.DictDataRepo{
		DB: db,
	}
	dictTypeRepo := &dict.DictTypeRepo{
		DB: db,
	}

	if err != nil {
		cleanup2()
		return nil, nil, err
	}
	menuRepo := &menu.MenuRepo{
		DB: db,
	}
	menuActionRepo := &menu.MenuActionRepo{
		DB: db,
	}
	loginSrv := &service.LoginSrv{

		UserRepo:       userRepo,
		UserRoleRepo:   userRoleRepo,
		RoleRepo:       roleRepo,
		RoleMenuRepo:   roleMenuRepo,
		MenuRepo:       menuRepo,
		MenuActionRepo: menuActionRepo,
	}

	trans := &util.Trans{
		DB: db,
	}
	menuSrv := &service.MenuSrv{
		TransRepo:              trans,
		MenuRepo:               menuRepo,
		MenuActionRepo:         menuActionRepo,
		MenuActionResourceRepo: menuActionResourceRepo,
	}

	userSrv := &service.UserSrv{
		TransRepo:    trans,
		UserRepo:     userRepo,
		UserRoleRepo: userRoleRepo,
		RoleRepo:     roleRepo,
	}

	dictSrv := &service.DictSrv{
		TransRepo:    trans,
		DictDataRepo: dictDataRepo,
		DictTypeRepo: dictTypeRepo,
	}

	injector := &InjectorForTalent{
		MenuSrv:  menuSrv,
		LoginSrv: loginSrv,
		DictSrv:  dictSrv,
		UserSrv:  userSrv,
	}
	return injector, func() {
		cleanup2()
	}, nil
}

func InitGormDB() (*gorm.DB, func(), error) {
	cfg := config.C.Gorm
	db, err := NewGormDB()
	if err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {}

	if cfg.EnableAutoMigrate {
		err = dao.AutoMigrate(db)
		if err != nil {
			return nil, cleanFunc, err
		}
	}

	return db, cleanFunc, nil
}

func NewGormDB() (*gorm.DB, error) {
	cfg := config.C
	var dsn string
	switch cfg.Gorm.DBType {
	case "mysql":
		dsn = cfg.MySQL.DSN()
	case "sqlite3":
		dsn = cfg.Sqlite3.DSN()
		_ = os.MkdirAll(filepath.Dir(dsn), 0777)
	case "postgres":
		dsn = cfg.Postgres.DSN()
	default:
		return nil, errors.New("unknown db")
	}

	return gormx.New(&gormx.Config{
		Debug:        cfg.Gorm.Debug,
		DBType:       cfg.Gorm.DBType,
		DSN:          dsn,
		MaxIdleConns: cfg.Gorm.MaxIdleConns,
		MaxLifetime:  cfg.Gorm.MaxLifetime,
		MaxOpenConns: cfg.Gorm.MaxOpenConns,
		TablePrefix:  cfg.Gorm.TablePrefix,
	})
}
