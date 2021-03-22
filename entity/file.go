package entity

import (
	"database/sql"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	URLFile sql.NullString `gorm:"size:50;unique;not null" json:"url_file"`
	UserID  int            `json:"user_id"`
}
