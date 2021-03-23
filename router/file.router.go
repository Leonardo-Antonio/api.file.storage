package router

import (
	"github.com/Leonardo-Antonio/api.file.store/handler"
	"github.com/Leonardo-Antonio/api.file.store/model"
	"github.com/gofiber/fiber/v2"
)

func User(storage model.IUser, app *fiber.App) {
	hand := handler.NewUser(storage)
	group := app.Group(BASE_URL)
	group.Post("/users/sign-in", hand.SignIn)
}