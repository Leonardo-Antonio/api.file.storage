package helper

import "database/sql"

func NullString(value string) sql.NullString {
	null := sql.NullString{String: value}
	if len(null.String) != 0 {
		null.Valid = true
	}
	return null
}
