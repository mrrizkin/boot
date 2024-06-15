package user

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/app/helpers"
	"github.com/mrrizkin/boot/app/models"
	"github.com/mrrizkin/boot/app/types"
	"github.com/mrrizkin/boot/system/logger"
)

type UserHandler struct {
	service *UserService
	log     *logger.Logger
}

type UserHandlerDeps struct {
	fx.In

	Service *UserService
	Logger  *logger.Logger
}

func NewUserHandler(p UserHandlerDeps) (*UserHandler, error) {
	return &UserHandler{
		service: p.Service,
		log:     p.Logger,
	}, nil
}

func (a *UserHandler) Create(c *fiber.Ctx) error {
	var (
		err     error
		payload models.User
	)

	err = c.BodyParser(&payload)
	if err != nil {
		return helpers.SendResponse(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	user, err := a.service.Create(&payload)
	if err != nil {
		return helpers.SendResponse(c, types.Response{
			Success: false,
			Message: "failed create user",
			Debug:   err.Error(),
		}, 500)
	}

	return helpers.SendResponse(c, types.Response{
		Success: true,
		Message: "success create user",
		Data:    user,
	})
}

func (a *UserHandler) FindAll(c *fiber.Ctx) error {
	var (
		err   error
		users []models.User
	)

	users, err = a.service.FindAll()
	if err != nil {
		return helpers.SendResponse(c, types.Response{
			Success: false,
			Message: "failed get users",
			Debug:   err.Error(),
		}, 500)
	}

	return helpers.SendResponse(c, types.Response{
		Success: true,
		Message: "success get users",
		Data:    users,
	})
}

func (a *UserHandler) FindByID(c *fiber.Ctx) error {
	var (
		err error
		id  int
	)

	id, err = c.ParamsInt("id")
	if err != nil {
		return helpers.SendResponse(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	user, err := a.service.FindByID(uint(id))
	if err != nil {
		return helpers.SendResponse(c, types.Response{
			Success: false,
			Message: "failed get user",
			Debug:   err.Error(),
		}, 500)
	}

	return helpers.SendResponse(c, types.Response{
		Success: true,
		Message: "success get user",
		Data:    user,
	})
}

func (a *UserHandler) Update(c *fiber.Ctx) error {
	var (
		err     error
		id      int
		payload models.User
	)

	id, err = c.ParamsInt("id")
	if err != nil {
		return helpers.SendResponse(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	err = c.BodyParser(&payload)
	if err != nil {
		return helpers.SendResponse(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	user, err := a.service.Update(uint(id), &payload)
	if err != nil {
		return helpers.SendResponse(c, types.Response{
			Success: false,
			Message: "failed update user",
			Debug:   err.Error(),
		}, 500)
	}

	return helpers.SendResponse(c, types.Response{
		Success: true,
		Message: "success update user",
		Data:    user,
	})
}

func (a *UserHandler) Delete(c *fiber.Ctx) error {
	var (
		err error
		id  int
	)

	id, err = c.ParamsInt("id")
	if err != nil {
		return helpers.SendResponse(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	err = a.service.Delete(uint(id))
	if err != nil {
		return helpers.SendResponse(c, types.Response{
			Success: false,
			Message: "failed delete user",
			Debug:   err.Error(),
		}, 500)
	}

	return helpers.SendResponse(c, types.Response{
		Success: true,
		Message: "success delete user",
	})
}
