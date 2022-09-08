package repository

import (
	"fmt"
	"github.com/wakataw/moku/entity"
	"gorm.io/gorm"
)

type programRepository struct {
	DB *gorm.DB
}

func (r *programRepository) Update(program *entity.Program) error {
	result := r.DB.Where("id = ?", program.ID).Updates(program)

	if result.RowsAffected == 0 {
		return ErrNoRowsAffected
	}

	return result.Error
}

func (r *programRepository) Delete(program *entity.Program) (err error) {
	err = r.DB.Unscoped().Delete(program).Error
	return
}

func (r *programRepository) FindById(programId int) (program *entity.Program, exists bool) {
	result := r.DB.Find(&program, programId)

	if result.RowsAffected == 0 {
		return program, false
	}

	return program, true
}

func (r *programRepository) All(lastCursor int, limit int, query string, ascending bool) (programs *[]entity.Program, err error) {
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

	err = tx.Find(&programs).Error

	if err != nil {
		return &[]entity.Program{}, err
	}

	return programs, nil
}

func (r *programRepository) Insert(program *entity.Program) (err error) {
	err = r.DB.Create(program).Error

	return
}

func NewProgramRepository(DB *gorm.DB) Program {
	return &programRepository{DB: DB}
}
