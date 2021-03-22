package helper

import "errors"

var (
	ErrRowsNotAffected = errors.New("there were no rows affected")
	ErrUserNotExist    = errors.New("username does not exist")
)
