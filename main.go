package main

import (
	"github.com/Leonardo-Antonio/api.file.store/DBUtil"
	"github.com/Leonardo-Antonio/api.file.store/app"
	"github.com/Leonardo-Antonio/api.file.store/entity"
	"log"
)

func main() {
	db := DBUtil.GetConnection(DBUtil.MARIADB)
	if err := db.AutoMigrate(&entity.User{}, &entity.File{}); err != nil {
		log.Fatalln(err)
	}

	server := app.New(app.DEV)
	server.Middlewares()
	server.Routers(db)
	server.Listening()
}
