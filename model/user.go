package model

import (
	"github.com/Leonardo-Antonio/api.file.store/entity"
	"github.com/Leonardo-Antonio/api.file.store/helper"
	"gorm.io/gorm"
)

type (
	user struct {
		db *gorm.DB
	}

	IUser interface {
		SignUp(user entity.User) error
		SignIn(credentials entity.User) (data entity.User, err error)
	}
)

func NewUser(db *gorm.DB) *user {
	return &user{db}
}

func (u *user) SignUp(user entity.User) error {
	result := u.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return helper.ErrRowsNotAffected
	}
	return nil
}

func (u *user) SignIn(credentials entity.User) (data entity.User, err error) {
	result := u.db.
		First(&data, "email = ? AND password = ?", credentials.Email, credentials.Password)

	if data.ID == 0 {
		return entity.User{}, helper.ErrUserNotExist
	}

	if result.Error != nil {
		return entity.User{}, err
	}

	return data, nil
}