// Package handlers provides HTTP request handlers and helper functions.
// This file contains utility functions used across multiple handlers
// to perform common tasks such as user retrieval, pagination, JSON response
// sending, and request body parsing and validation.
package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/boot/app/models"
	"github.com/mrrizkin/boot/system/types"
)

// getUser retrieves the user from the database based on the user ID stored in the request context.
// Returns nil if the user is not found or an error occurs.
func (h *Handlers) getUser(c *fiber.Ctx) *models.User {
	userId := c.Locals("uid").(uint)
	user, err := h.userRepo.FindByID(userId)
	if err != nil {
		return nil
	}

	return user
}

// getPaginateQuery extracts pagination parameters from the request query.
// It returns a Pagination struct with default values if not specified in the query.
func (h *Handlers) getPaginateQuery(c *fiber.Ctx) types.Pagination {
	return types.Pagination{
		Page:    c.QueryInt("page", 1),
		PerPage: c.QueryInt("per_page", 10),
	}
}

// sendJson sends a JSON response with the given data and status code.
// If no status code is provided, it defaults to 200 (OK).
func (h *Handlers) sendJson(c *fiber.Ctx, resp interface{}, status ...int) error {
	var statusCode int

	if len(status) == 0 {
		statusCode = 200
	} else {
		statusCode = status[0]
	}

	return c.Status(statusCode).JSON(resp)
}

// bodyParseValidate parses the request body into the given struct and validates it.
// It returns an error if parsing fails or if the data doesn't pass validation.
func (h *Handlers) bodyParseValidate(c *fiber.Ctx, out interface{}) error {
	err := c.BodyParser(out)
	if err != nil {
		h.Log("error", "failed to parse payload", "err", err)
		return &fiber.Error{
			Code:    400,
			Message: "payload not valid",
		}
	}

	validationError := h.System.Validator.MustValidate(out)
	if validationError != nil {
		return validationError
	}

	return nil
}
