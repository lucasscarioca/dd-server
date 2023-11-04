package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lucasscarioca/dinodiary/pkg/echolambda"
)

type Router struct {
	*echo.Echo
}

func NewRouter(
	authMiddleware func() echo.MiddlewareFunc,
	userHandler UserHandler,
	authHandler AuthHandler,
) (*Router, error) {
	e := echo.New()

	e.Use(customLogger())
	e.Use(middleware.Recover())
	e.Use(customCORS())

	// Docs
	// e.GET("/docs/*any", swaggerWrapper)

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	v1 := e.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/forgot", authHandler.Forgot)
			auth.PUT("/reset/:token", authHandler.Reset)
		}
		user := v1.Group("/users", authMiddleware())
		{
			user.GET("/", userHandler.List)
		}
	}

	return &Router{
		e,
	}, nil
}

func (r *Router) Serve(listenAddr string) {
	r.Logger.Fatal(r.Start(listenAddr))
}

func (r *Router) ServeLambda() {
	lambdaAdapter := &echolambda.LambdaAdapter{Echo: r.Echo}
	lambda.Start(lambdaAdapter.Handler)
}
