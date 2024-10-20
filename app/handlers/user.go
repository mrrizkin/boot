package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/boot/app/models"
	"github.com/mrrizkin/boot/system/types"
	_ "github.com/mrrizkin/boot/system/validator"
)

// UserCreate godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user with the provided information
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.User							true	"User information"
//	@Success		200		{object}	types.Response{data=models.User}	"Successfully created user"
//	@Failure		400		{object}	validator.GlobalErrorResponse		"Bad request"
//	@Failure		500		{object}	validator.GlobalErrorResponse		"Internal server error"
//	@Router			/user [post]
func (h *Handlers) UserCreate(c *fiber.Ctx) error {
	payload := new(models.User)
	err := h.bodyParseValidate(c, payload)
	if err != nil {
		return err
	}

	user, err := h.userService.Create(payload)
	if err != nil {
		h.Log("error", "failed to create user", "err", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("failed to create user: %s", err),
		}
	}

	return h.sendJson(c, types.Response{
		Status:  "success",
		Message: "user created successfully",
		Data:    user,
	})
}

// UserFindAll godoc
//
//	@Summary		Get all users
//	@Description	Retrieve a list of all users with pagination
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int																false	"Page number"
//	@Param			per_page	query		int																false	"Number of items per page"
//	@Success		200			{object}	types.Response{data=[]models.User,meta=types.PaginationMeta}	"Successfully retrieved users"
//	@Failure		500			{object}	validator.GlobalErrorResponse									"Internal server error"
//	@Router			/user [get]
func (h *Handlers) UserFindAll(c *fiber.Ctx) error {
	pagination := h.getPaginateQuery(c)
	users, err := h.userService.FindAll(pagination)
	if err != nil {
		h.Log("error", "failed to get users", "err", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("failed to get users: %s", err),
		}
	}

	return h.sendJson(c, types.Response{
		Status:  "success",
		Message: "users retrieved successfully",
		Data:    users.Result,
		Meta: &types.PaginationMeta{
			Page:      pagination.Page,
			PerPage:   pagination.PerPage,
			Total:     users.Total,
			PageCount: users.Total / pagination.PerPage,
		},
	})
}

// UserFindByID godoc
//
//	@Summary		Get a user by ID
//	@Description	Retrieve a user by their ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int									true	"User ID"
//	@Success		200	{object}	types.Response{data=models.User}	"Successfully retrieved user"
//	@Failure		400	{object}	validator.GlobalErrorResponse		"Bad request"
//	@Failure		404	{object}	validator.GlobalErrorResponse		"User not found"
//	@Failure		500	{object}	validator.GlobalErrorResponse		"Internal server error"
//	@Router			/user/{id} [get]
func (h *Handlers) UserFindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.Log("error", "failed to parse id", "err", err)
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "invalid id",
		}
	}

	user, err := h.userService.FindByID(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			return &fiber.Error{
				Code:    fiber.StatusNotFound,
				Message: "user not found",
			}
		}

		h.Log("error", "failed to get user", "err", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("failed to get user: %s", err),
		}
	}

	return h.sendJson(c, types.Response{
		Status:  "success",
		Message: "user retrieved successfully",
		Data:    user,
	})
}

// UserUpdate godoc
//
//	@Summary		Update a user
//	@Description	Update a user's information by their ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int									true	"User ID"
//	@Param			user	body		models.User							true	"Updated user information"
//	@Success		200		{object}	types.Response{data=models.User}	"Successfully updated user"
//	@Failure		400		{object}	validator.GlobalErrorResponse		"Bad request"
//	@Failure		500		{object}	validator.GlobalErrorResponse		"Internal server error"
//	@Router			/user/{id} [put]
func (h *Handlers) UserUpdate(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.Log("error", "failed to parse id", "err", err)
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "invalid id",
		}
	}

	payload := new(models.User)
	err = h.bodyParseValidate(c, payload)
	if err != nil {
		return err
	}

	user, err := h.userService.Update(uint(id), payload)
	if err != nil {
		h.Log("error", "failed to update user", "err", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("failed to update user: %s", err),
		}
	}

	return h.sendJson(c, types.Response{
		Status:  "success",
		Message: "user updated successfully",
		Data:    user,
	})
}

// UserDelete godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user by their ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int								true	"User ID"
//	@Success		200	{object}	types.Response					"Successfully deleted user"
//	@Failure		400	{object}	validator.GlobalErrorResponse	"Bad request"
//	@Failure		401	{object}	validator.GlobalErrorResponse	"Unauthorized"
//	@Failure		500	{object}	validator.GlobalErrorResponse	"Internal server error"
//	@Router			/user/{id} [delete]
func (h *Handlers) UserDelete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.Log("error", "failed to parse id", "err", err)
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "invalid id",
		}
	}

	user := h.getUser(c)
	if user == nil {
		return &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized",
		}
	}

	err = h.userService.Delete(user, uint(id))
	if err != nil {
		h.Log("error", "failed to delete user", "err", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("failed to delete user: %s", err),
		}
	}

	return h.sendJson(c, types.Response{
		Status:  "success",
		Message: "user deleted successfully",
	})
}
