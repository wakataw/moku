package repository

import "github.com/wakataw/moku/entity"

type PermissionRepository interface {
	Insert(perm *entity.Permission) (*entity.Permission, error)
	Update(perm *entity.Permission) (*entity.Permission, error)
	Delete(perm *entity.Permission) (err error)
	FindById(permId int) (permission *entity.Permission)
	FindByIds(permIds ...int) (permissions *[]entity.Permission)
	FindByName(permName string) (permission *entity.Permission, exists bool)
	FindByNames(permNames ...string) (permissions *[]entity.Permission)
	All(lastCursor int, limit int, query string, ascending bool) (permissions *[]entity.Permission, err error)
}
