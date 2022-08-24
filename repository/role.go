package repository

import "github.com/wakataw/moku/entity"

type RoleRepository interface {
	Insert(role *entity.Role) (*entity.Role, error)
	Update(role *entity.Role) (*entity.Role, error)
	Delete(role *entity.Role) (err error)
	FindById(roleId int) (role *entity.Role)
	FindByIds(roleIds ...int) (roles *[]entity.Role)
	FindByName(roleName string) (role *entity.Role, exists bool)
	FindByNames(roleNames ...string) (roles *[]entity.Role)
	FindPermissions(roleId int) (permissions *[]entity.Permission)
	HasPermission(roleId int, permissionName string) bool
	All(lastCursor int, limit int, query string, ascending bool) (roles *[]entity.Role, err error)
}
