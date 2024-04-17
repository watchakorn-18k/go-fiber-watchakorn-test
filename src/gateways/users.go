package gateways

import (
	"go-fiber-proton/domain/entities"

	"github.com/gofiber/fiber/v2"
)

func (h HTTPGateway) GetAllUserData(ctx *fiber.Ctx) error {

	data, err := h.UserService.GetAllUser()
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot get all users data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}

func (h HTTPGateway) CreateNewUserAccount(ctx *fiber.Ctx) error {

	var bodyData entities.NewUserBody
	if err := ctx.BodyParser(&bodyData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	if bodyData == (entities.NewUserBody{}) {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	status := h.UserService.InsertNewAccount(&bodyData)

	if !status {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot insert new user account."})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}

func (h HTTPGateway) UpdateUser(ctx *fiber.Ctx) error {
	var bodyData entities.NewUserBody
	if err := ctx.BodyParser(&bodyData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	if bodyData == (entities.NewUserBody{}) {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	if !h.UserService.UpdateUser(&bodyData) {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot update user account."})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "update success"})
}

func (h HTTPGateway) DeleteUser(ctx *fiber.Ctx) error {
	paramAll := ctx.Queries()
	name := paramAll["name"]
	if name == "" {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "query param 'name' cannot be empty."})
	}
	if !h.UserService.DeleteUser(name) {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot delete user account."})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "delete success"})
}
