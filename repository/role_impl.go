package repository

import (
	"fmt"
	"github.com/wakataw/moku/entity"
	"gorm.io/gorm"
)

type roleRepository struct {
	DB *gorm.DB
}

func (r roleRepository) All(lastCursor int, limit int, query string, ascending bool) (roles *[]entity.Role, err error) {

	tx := r.DB.Where("name like ?", fmt.Sprintf("%%%v%%", query))

	// pagination
	if lastCursor > 0 {
		if ascending {
			tx.Where("id > ?", lastCursor)
		} else {
			tx.Where("id < ?", lastCursor)
		}
	}

	// order
	if ascending {
		tx.Order("id asc")
	} else {
		tx.Order("id desc")
	}

	// add limit
	tx.Limit(limit)

	err = tx.Find(&roles).Error

	if err != nil {
		return &[]entity.Role{}, err
	}

	return roles, nil
}

func (r roleRepository) Insert(role *entity.Role) (*entity.Role, error) {
	result := r.DB.Create(role)

	return role, result.Error
}

func (r roleRepository) Update(role *entity.Role) (*entity.Role, error) {
	result := r.DB.Select("name").Save(role)
	return role, result.Error
}

func (r roleRepository) Delete(role *entity.Role) (err error) {
	result := r.DB.Delete(role)
	err = result.Error
	return
}

func (r roleRepository) FindById(roleId int) (role *entity.Role) {
	r.DB.Find(role, roleId)
	return
}

func (r roleRepository) FindByIds(roleIds ...int) (roles *[]entity.Role) {
	r.DB.Find(roles, roleIds)
	return
}

func (r roleRepository) FindByName(roleName string) (role *entity.Role, exists bool) {
	result := r.DB.Preload("Permissions").Where("name = ?", roleName).Find(&role)

	if result.RowsAffected == 0 {
		exists = false
	} else {
		exists = true
	}

	return
}

func (r roleRepository) FindByNames(roleNames ...string) (roles *[]entity.Role) {
	r.DB.Where("name IN ?", roleNames).Find(roles)
	return
}

func (r roleRepository) FindPermissions(roleId int) (permissions *[]entity.Permission) {
	var role entity.Role
	r.DB.Preload("Permissions").Find(&role, roleId)

	return &role.Permissions
}

func (r roleRepository) HasPermission(roleId int, permissionName string) bool {
	var role entity.Role
	result := r.DB.Where("id = ? AND permissions.name = ?", roleId, permissionName).Find(&role)

	if result.RowsAffected >= 0 {
		return true
	}

	return false
}

func NewRoleRepository(DB *gorm.DB) RoleRepository {
	return &roleRepository{DB: DB}
}
