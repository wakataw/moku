package repository

import (
	"github.com/wakataw/moku/entity"
	"gorm.io/gorm"
)

type permissionRepository struct {
	DB *gorm.DB
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
