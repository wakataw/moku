package service

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/wakataw/moku/entity"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/pkg"
	"github.com/wakataw/moku/repository"
)

type programService struct {
	repository repository.Program
}

func (p programService) Update(request *model.UpdateProgramRequest) (response *model.GetProgramResponse, err error) {
	var program entity.Program
	program.ID = request.ID
	program.Name = request.Name
	program.Description = request.Description
	program.Start = request.Start
	program.End = request.End
	program.Public = request.Public
	program.Show = request.Show
	program.UpdatedBy = &request.UpdatedBy

	err = p.repository.Update(&program)

	if err != nil {
		return nil, err
	}

	return &model.GetProgramResponse{
		ID:          program.ID,
		Name:        program.Name,
		Description: program.Description,
		Start:       program.Start,
		End:         program.End,
		Show:        program.Show,
		Public:      program.Public,
		CreatedUser: nil,
		UpdatedUser: nil,
	}, nil

}

func (p programService) Delete(programId int) (err error) {
	err = p.repository.Delete(&entity.Program{ID: programId})
	return
}

func (p programService) GetProgramById(programId int) (response *model.GetProgramResponse, err error) {
	program, exists := p.repository.FindById(programId)

	if !exists {
		return nil, ErrObjectDoesntExists
	}

	return &model.GetProgramResponse{
		ID:          program.ID,
		Name:        program.Name,
		Description: program.Description,
		Start:       program.Start,
		End:         program.End,
		Show:        program.Show,
		Public:      program.Public,
		CreatedUser: &model.GetUserResponseSimple{
			ID:       program.CreatedUser.ID,
			Username: program.CreatedUser.Username,
			FullName: program.CreatedUser.FullName,
		},
		UpdatedUser: &model.GetUserResponseSimple{
			ID:       program.UpdatedUser.ID,
			Username: program.UpdatedUser.Username,
			FullName: program.UpdatedUser.FullName,
		},
	}, nil

}

func (p programService) Create(request *model.CreateProgramRequest) (response *model.GetProgramResponse, err error) {
	var program entity.Program

	program.Name = request.Name
	program.Description = request.Description
	program.Slug = fmt.Sprintf("%v-%v", slug.Make(request.Name), pkg.NewLen(8))
	program.Start = request.Start
	program.End = request.End
	program.Show = request.Show
	program.Public = request.Public
	program.CreatedBy = &request.CreatedBy

	err = p.repository.Insert(&program)

	if err != nil {
		return nil, err
	}

	return &model.GetProgramResponse{
		ID:          program.ID,
		Name:        program.Name,
		Description: program.Description,
		Start:       program.Start,
		End:         program.End,
		Show:        program.Show,
		Public:      program.Public,
	}, nil

}

func (p programService) All(request *model.RequestParameter) (responses *[]model.GetProgramResponse, err error) {
	programs, err := p.repository.All(*request.LastCursor, request.Limit, request.Query, request.Ascending)

	if err != nil {
		return &[]model.GetProgramResponse{}, err
	}

	var programResponse []model.GetProgramResponse

	for _, v := range *programs {
		programResponse = append(programResponse, model.GetProgramResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Start:       v.Start,
			End:         v.End,
			Show:        v.Show,
			Public:      v.Public,
			CreatedUser: &model.GetUserResponseSimple{
				ID:       v.CreatedUser.ID,
				Username: v.CreatedUser.Username,
				FullName: v.CreatedUser.FullName,
			},
			UpdatedUser: &model.GetUserResponseSimple{
				ID:       v.UpdatedUser.ID,
				Username: v.UpdatedUser.Username,
				FullName: v.UpdatedUser.FullName,
			},
		})
	}

	return &programResponse, nil
}

func NewProgramService(repository *repository.Program) ProgramService {
	return &programService{repository: *repository}
}
