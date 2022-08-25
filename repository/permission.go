package repository

import "github.com/wakataw/moku/entity"

type PermissionRepository interface {
	Insert(perm *entity.Permission) (*entity.Permission, error)
	Update(perm *entity.Permission) error
	Delete(perm *entity.Permission) error
	FindById(permId int) *entity.Permission
	FindByIds(permIds ...int) *[]entity.Permission
	FindByName(permName string) *entity.Permission
	FindByNames(permNames ...string) *[]entity.Permission
}
