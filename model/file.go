package model

import (
	"github.com/Leonardo-Antonio/api.file.store/entity"
	"github.com/Leonardo-Antonio/api.file.store/helper"
	"gorm.io/gorm"
)

type (
	file struct {
		db *gorm.DB
	}
	IFile interface {
	}
)

func NewFile(db *gorm.DB) *file {
	return &file{db}
}

func (f file) Create(file entity.File) error {
	result := f.db.Create(&file)
	if result.RowsAffected != 1 {
		return helper.ErrRowsNotAffected
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
