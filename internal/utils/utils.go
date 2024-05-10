package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func SendResponse(c *fiber.Ctx, resp interface{}, status ...int) error {
	var statusCode int

	if len(status) == 0 {
		statusCode = 200
	} else {
		statusCode = status[0]
	}

	return c.Status(statusCode).JSON(resp)
}

func EnvStr(name string, def ...string) (*string, error) {
	env := os.Getenv(name)

	if env == "" {
		for _, v := range def {
			return &v, nil
		}

		return nil, fmt.Errorf("env %s is empty", name)
	}

	return &env, nil
}

func EnvBool(name string, def ...bool) (*bool, error) {
	env := os.Getenv(name)

	if env == "" {
		for _, v := range def {
			return &v, nil
		}

		return nil, fmt.Errorf("env %s is empty", name)
	}

	value, err := strconv.ParseBool(env)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func EnvInt(name string, def ...int) (*int, error) {
	env := os.Getenv(name)

	if env == "" {
		for _, v := range def {
			return &v, nil
		}

		return nil, fmt.Errorf("env %s is empty", name)
	}

	value, err := strconv.Atoi(env)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, option := range options {
		option(componentHandler)
	}

	return adaptor.HTTPHandler(componentHandler)(c)
}
