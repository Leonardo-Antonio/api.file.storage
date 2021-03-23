package handler

import (
	"encoding/json"
	"github.com/Leonardo-Antonio/api.file.store/entity"
	"github.com/Leonardo-Antonio/api.file.store/helper"
	"github.com/Leonardo-Antonio/api.file.store/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type user struct {
	storage model.IUser
}

func NewUser(storage model.IUser) *user {
	return &user{storage}
}

func (u *user) SignIn(c *fiber.Ctx) error {
	var credential entity.User
	err := json.Unmarshal(c.Body(), &credential)
	if err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(helper.Response{MessageType: helper.ERR, Message: err.Error(), Error: true, Data: nil})
	}

	// data, err := u.storage.SignIn(credential)
	//if err != nil {
	//	return c.Status(http.StatusInternalServerError).
	//		JSON(helper.Response{MessageType: helper.ERR, Message: err.Error(), Error: true, Data: nil})
	//}

	return c.Status(http.StatusOK).
		JSON(helper.Response{MessageType: helper.MSG, Message: "record created successfully", Error: false, Data: "data"})
}