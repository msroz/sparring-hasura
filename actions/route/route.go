package route

import (
	"net/http"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/msroz/sparring-hasura/actions/graph"
	"github.com/msroz/sparring-hasura/actions/graph/generated"
	"github.com/msroz/sparring-hasura/actions/handler"
)

func Init() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi!")
	})

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

	/* GraphQL Server with gqlgen */
	srv := gqlHandler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{}},
		),
	)
	srv.AddTransport(transport.POST{})

	// https://github.com/99designs/gqlgen/blob/master/example/fileupload/server/server.go
	var mb int64 = 1 << 20
	srv.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * mb,
		MaxUploadSize: 50 * mb,
	})

	// See: https://gqlgen.com/reference/complexity/
	// graphqlHandler.Use(extension.FixedComplexityLimit(10))
	e.POST("/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	playgroundHandler := playground.Handler("GraphQL", "/query")
	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	return e
}
