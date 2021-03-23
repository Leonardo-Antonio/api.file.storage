package entity

import (
	"database/sql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name sql.NullString `gorm:"size:100;not null" json:"name"`
	LastName sql.NullString `gorm:"size:100;not null" json:"last_name"`
	Email sql.NullString `gorm:"size:150;unique;not null" json:"email"`
	Password sql.NullString `gorm:"size:100;not null" json:"password"`
	Files []File `json:"files"`
}
