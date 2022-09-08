package service

import "github.com/wakataw/moku/model"

type ProgramService interface {
	Create(request *model.CreateProgramRequest) (response *model.GetProgramResponse, err error)
	All(request *model.RequestParameter) (responses *[]model.GetProgramResponse, err error)
	GetProgramById(programId int) (response *model.GetProgramResponse, err error)
	Delete(programId int) (err error)
	Update(request *model.UpdateProgramRequest) (response *model.GetProgramResponse, err error)
}
