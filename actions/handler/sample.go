package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Actions
func QuerySample() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"hello": "world"})
	}
}

type MutationSampleParam struct {
	Input struct {
		Message string `json:"message"`
	} `json:"input"`
}

func MutationSample() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(MutationSampleParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"hello": param.Input.Message})
	}
}

type TodoParam struct {
	ID string `json:"id"`
}

// Events
func CreateTodoSample() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(TodoParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		fmt.Println(param.ID)
		return c.String(http.StatusOK, "Hello, World!")
	}
}
