package service

import (
	"github.com/wakataw/moku/entity"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/repository"
)

type permissionService struct {
	repository repository.PermissionRepository
}

func (p permissionService) Create(request *model.CreatePermissionRequest) (response *model.GetPermissionResponse, err error) {
	var perm *entity.Permission

	if perm, err = p.repository.Insert(&entity.Permission{
		Name: request.Name,
	}); err != nil {
		return response, err
	}

	return &model.GetPermissionResponse{
		ID:   perm.ID,
		Name: perm.Name,
	}, nil
}

func NewPermissionService(repository *repository.PermissionRepository) PermissionService {
	return &permissionService{repository: *repository}
}
