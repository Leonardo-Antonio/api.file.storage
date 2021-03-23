package app

import (
	"github.com/Leonardo-Antonio/api.file.store/model"
	"github.com/Leonardo-Antonio/api.file.store/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
	"log"
	"os"
)

type app struct {
	app   *fiber.App
	State string
}

func New(state string) *app {
	return &app{
		app:   fiber.New(),
		State: state,
	}
}

func (server *app) Middlewares() {
	server.app.Use(logger.New())
}



func (server *app) Routers(db *gorm.DB) {
	storageUser := model.NewUser(db)
	router.User(storageUser, server.app)
}

func (server *app) Listening() {
	switch server.State {
	case DEV:
		log.Fatalln(server.app.Listen(":8080"))
	case PROD:
		log.Fatalln(server.app.Listen(":" + os.Getenv("PORT")))
	default:
		log.Fatalln("enter a valid state [app.PROD || app.DEV]")
	}
}
