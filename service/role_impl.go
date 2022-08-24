package service

import (
	"errors"
	"github.com/wakataw/moku/entity"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/repository"
)

type roleService struct {
	repository repository.RoleRepository
}

func (r roleService) All(request *model.RequestParameter) (responses *[]model.GetRoleSimpleResponse, err error) {
	results, err := r.repository.All(
		*request.LastCursor,
		request.Limit,
		request.Query,
		request.Ascending,
	)

	if err != nil {
		return &[]model.GetRoleSimpleResponse{}, err
	}

	var roleResp []model.GetRoleSimpleResponse

	for _, v := range *results {
		roleResp = append(roleResp, model.GetRoleSimpleResponse{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return &roleResp, nil
}

func (r roleService) Create(request *model.CreateRoleRequest) (response *model.GetRoleResponse, err error) {
	role, exists := r.repository.FindByName(request.Name)

	if !exists {
		role, err = r.repository.Insert(&entity.Role{
			Name: request.Name,
		})

		if err != nil {
			return nil, err
		}
	}

	var perms []model.GetPermissionResponse

	for _, v := range role.Permissions {
		perms = append(perms, model.GetPermissionResponse{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return &model.GetRoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Permissions: perms,
	}, nil

}

func (r roleService) Update(request *model.UpdateRoleRequest) (response *model.GetRoleResponse, err error) {
	var permissions []entity.Permission

	for _, v := range request.Permissions {
		permissions = append(permissions, entity.Permission{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	role, err := r.repository.Update(&entity.Role{
		ID:          request.ID,
		Name:        request.Name,
		Permissions: permissions,
	})

	if err != nil {
		return &model.GetRoleResponse{}, err
	}

	var permissionsResp []model.GetPermissionResponse

	for _, v := range role.Permissions {
		permissionsResp = append(permissionsResp, model.GetPermissionResponse{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return &model.GetRoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Permissions: permissionsResp,
	}, nil

}

func (r roleService) Delete(roleId int) (err error) {
	err = r.repository.Delete(&entity.Role{ID: roleId})

	return err
}

func (r roleService) GetRoleById(roleId int) *model.GetRoleResponse {
	//TODO implement me
	panic("implement me")
}

func (r roleService) GetRoleByName(roleName string) (*model.GetRoleResponse, error) {
	role, exists := r.repository.FindByName(roleName)

	if !exists {
		return &model.GetRoleResponse{}, errors.New("role doesn't exists")
	}
	var permissions []model.GetPermissionResponse

	for _, v := range role.Permissions {
		permissions = append(permissions, model.GetPermissionResponse{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return &model.GetRoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Permissions: permissions,
	}, nil
}

func NewRoleService(roleRepository *repository.RoleRepository) RoleService {
	return &roleService{
		*roleRepository,
	}
}
