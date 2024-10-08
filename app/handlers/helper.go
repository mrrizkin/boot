package handlers

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/boot/app/models"
	"github.com/mrrizkin/boot/system/stypes"
	"github.com/mrrizkin/boot/system/view/tag"
)

type StreamResponse struct {
	ID    string
	Event string
	Data  interface{}
}

func (h *Handlers) GetUser(c *fiber.Ctx) *models.User {
	userId := c.Locals("uid").(uint)
	user, err := h.userRepo.FindByID(userId)
	if err != nil {
		return nil
	}

	return user
}

func (h *Handlers) SendJson(c *fiber.Ctx, resp interface{}, status ...int) error {
	var statusCode int

	if len(status) == 0 {
		statusCode = 200
	} else {
		statusCode = status[0]
	}

	return c.Status(statusCode).JSON(resp)
}

func (h *Handlers) SendStream(w *bufio.Writer, resp *StreamResponse) error {
	_, err := fmt.Fprintf(w, "id: %s\nevent: %s\ndata: %s\n\n", resp.ID, resp.Event, resp.Data)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handlers) GetPaginationQuery(c *fiber.Ctx) stypes.Pagination {
	page := c.QueryInt("page", 1)
	perPage := c.QueryInt("per_page", 10)

	return stypes.Pagination{
		Page:    page,
		PerPage: perPage,
	}
}

func (h *Handlers) Render(c *fiber.Ctx, name string, data ...fiber.Map) error {
	ctx := fiber.Map{}
	if len(data) > 0 {
		ctx = data[0]
	}

	id, err := tag.State.GenerateID()
	if err != nil {
		return err
	}
	ctx["gonja-tag-state-id"] = id
	var buf bytes.Buffer
	err = h.System.View.Render(&buf, name, ctx)
	if err != nil {
		return err
	}

	tag.State.Clear(id)

	return c.Type("html").Send(buf.Bytes())
}
