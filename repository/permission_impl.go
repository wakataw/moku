package repository

import (
	"fmt"
	"github.com/wakataw/moku/entity"
	"gorm.io/gorm"
)

type permissionRepository struct {
	DB *gorm.DB
}

func (r permissionRepository) All(lastCursor int, limit int, query string, ascending bool) (permissions *[]entity.Permission, err error) {
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

	err = tx.Find(&permissions).Error

	if err != nil {
		return &[]entity.Permission{}, err
	}

	return permissions, nil
}

func (r permissionRepository) Insert(perm *entity.Permission) (*entity.Permission, error) {
	result := r.DB.Where("name = ?", perm.Name).Find(&perm)

	if result.Error != nil {
		return &entity.Permission{}, result.Error
	}

	if result.RowsAffected > 0 {
		return perm, nil
	}

	// insert new permission
	result = r.DB.Create(&perm)

	if result.Error != nil {
		return &entity.Permission{}, nil
	}

	return perm, nil

}

func (r permissionRepository) Update(perm *entity.Permission) error {
	//TODO implement me
	panic("implement me")
}

func (r permissionRepository) Delete(perm *entity.Permission) error {
	//TODO implement me
	panic("implement me")
}

func (r permissionRepository) FindById(permId int) *entity.Permission {
	//TODO implement me
	panic("implement me")
}

func (r permissionRepository) FindByIds(permIds ...int) *[]entity.Permission {
	//TODO implement me
	panic("implement me")
}

func (r permissionRepository) FindByName(permName string) *entity.Permission {
	//TODO implement me
	panic("implement me")
}

func (r permissionRepository) FindByNames(permNames ...string) *[]entity.Permission {
	//TODO implement me
	panic("implement me")
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{
		DB: db,
	}
}
