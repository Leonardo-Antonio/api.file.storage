package DBUtil

import (
	"gorm.io/gorm"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetConnection(typeConnection string) *gorm.DB {
	once.Do(func() {
		con := newConnection()
		switch typeConnection {
		case MARIADB:
			db = con.mariaDB()
		case MYSQL:
			db = con.mysql()
		case MSSQL:
			db = con.mssql()
		case PSSQL:
			db = con.pssql()
		}
	})
	return db
}
