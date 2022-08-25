package service

import (
	"errors"
	"fmt"
	"github.com/wakataw/moku/entity"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/repository"
)

type permissionService struct {
	repository repository.PermissionRepository
}

func (p *permissionService) Update(request *model.UpdatePermissionRequest) (response *model.GetPermissionResponse, err error) {
	perm, err := p.repository.Update(&entity.Permission{
		ID:   request.ID,
		Name: request.Name,
	})

	if err != nil {
		return &model.GetPermissionResponse{}, err
	}

	return &model.GetPermissionResponse{
		ID:   perm.ID,
		Name: perm.Name,
	}, nil
}

func (p *permissionService) Delete(permId int) (err error) {
	err = p.repository.Delete(&entity.Permission{ID: permId})
	return
}

func (p *permissionService) GetPermissionById(permId int) *model.GetPermissionResponse {
	panic("implement me")
}

func (p *permissionService) GetPermissionByName(permName string) (*model.GetPermissionResponse, error) {
	perm, exists := p.repository.FindByName(permName)

	if !exists {
		return &model.GetPermissionResponse{}, errors.New(fmt.Sprintf("permission %v doesn't exist", permName))
	}

	return &model.GetPermissionResponse{
		ID:   perm.ID,
		Name: perm.Name,
	}, nil

}

func (p *permissionService) All(request *model.RequestParameter) (responses *[]model.GetPermissionResponse, err error) {
	results, err := p.repository.All(
		*request.LastCursor,
		request.Limit,
		request.Query,
		request.Ascending,
	)

	if err != nil {
		return &[]model.GetPermissionResponse{}, err
	}

	var permResp []model.GetPermissionResponse

	for _, v := range *results {
		permResp = append(permResp, model.GetPermissionResponse{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return &permResp, nil
}

func (p *permissionService) Create(request *model.CreatePermissionRequest) (response *model.GetPermissionResponse, err error) {
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
