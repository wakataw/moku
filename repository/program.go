package repository

import "github.com/wakataw/moku/entity"

type Program interface {
	All(lastCursor int, limit int, query string, ascending bool) (programs *[]entity.Program, err error)
	Insert(program *entity.Program) error
	Update(program *entity.Program) error
	Delete(program *entity.Program) (err error)
	FindById(programId int) (program *entity.Program, exists bool)
}
