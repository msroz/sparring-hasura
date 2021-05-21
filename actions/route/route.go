package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/msroz/sparring-hasura/actions/handler"
)

func Init() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	actions := e.Group("/actions")
	{
		actions.POST("/querySample", handler.QuerySample())
		actions.POST("/mutationSample", handler.MutationSample())

		actions.POST("/registerUser", handler.RegisterUser())
	}

	events := e.Group("/events")
	{
		events.POST("/createTodo", handler.CreateTodoSample())
	}

	return e
}
