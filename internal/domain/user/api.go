package user

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/internal/model"
	"github.com/mrrizkin/boot/internal/system/logger"
	"github.com/mrrizkin/boot/internal/types"
	"github.com/mrrizkin/boot/internal/utils"
)

type UserAPI struct {
	service *UserService
	log     *logger.Logger
}

type UserAPIParams struct {
	fx.In

	Service *UserService
	Logger  *logger.Logger
}

func NewUserAPI(p UserAPIParams) (*UserAPI, error) {
	return &UserAPI{
		service: p.Service,
		log:     p.Logger,
	}, nil
}

func (a *UserAPI) Create(c *fiber.Ctx) error {
	var (
		err     error
		payload model.User
	)

	err = c.BodyParser(&payload)
	if err != nil {
		return utils.SendResponse(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	user, err := a.service.Create(&payload)
	if err != nil {
		return utils.SendResponse(c, types.Response{
			Success: false,
			Message: "failed create user",
			Debug:   err.Error(),
		}, 500)
	}

	return utils.SendResponse(c, types.Response{
		Success: true,
		Message: "success create user",
		Data:    user,
	})
}

func (a *UserAPI) FindAll(c *fiber.Ctx) error {
	var (
		err   error
		users []model.User
	)

	users, err = a.service.FindAll()
	if err != nil {
		return utils.SendResponse(c, types.Response{
			Success: false,
			Message: "failed get users",
			Debug:   err.Error(),
		}, 500)
	}

	return utils.SendResponse(c, types.Response{
		Success: true,
		Message: "success get users",
		Data:    users,
	})
}

func (a *UserAPI) FindByID(c *fiber.Ctx) error {
	var (
		err error
		id  int
	)

	id, err = c.ParamsInt("id")
	if err != nil {
		return utils.SendResponse(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	user, err := a.service.FindByID(uint(id))
	if err != nil {
		return utils.SendResponse(c, types.Response{
			Success: false,
			Message: "failed get user",
			Debug:   err.Error(),
		}, 500)
	}

	return utils.SendResponse(c, types.Response{
		Success: true,
		Message: "success get user",
		Data:    user,
	})
}

func (a *UserAPI) Update(c *fiber.Ctx) error {
	var (
		err     error
		id      int
		payload model.User
	)

	id, err = c.ParamsInt("id")
	if err != nil {
		return utils.SendResponse(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	err = c.BodyParser(&payload)
	if err != nil {
		return utils.SendResponse(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	user, err := a.service.Update(uint(id), &payload)
	if err != nil {
		return utils.SendResponse(c, types.Response{
			Success: false,
			Message: "failed update user",
			Debug:   err.Error(),
		}, 500)
	}

	return utils.SendResponse(c, types.Response{
		Success: true,
		Message: "success update user",
		Data:    user,
	})
}

func (a *UserAPI) Delete(c *fiber.Ctx) error {
	var (
		err error
		id  int
	)

	id, err = c.ParamsInt("id")
	if err != nil {
		return utils.SendResponse(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	err = a.service.Delete(uint(id))
	if err != nil {
		return utils.SendResponse(c, types.Response{
			Success: false,
			Message: "failed delete user",
			Debug:   err.Error(),
		}, 500)
	}

	return utils.SendResponse(c, types.Response{
		Success: true,
		Message: "success delete user",
	})
}
