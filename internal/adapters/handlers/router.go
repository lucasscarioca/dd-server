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
	assistHandler AssistHandler,
	authHandler AuthHandler,
	entryHandler EntryHandler,
	dinoHandler DinoHandler,
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
			user.GET("/profile", userHandler.Profile)
			user.GET("/:id", userHandler.Find)
			user.PUT("/", userHandler.Update)
			user.DELETE("/", userHandler.Delete)

			// Assist
			user.GET("/assisted/requests", assistHandler.ListAssistedUsersRequests)
			user.GET("/assisted", assistHandler.ListAssistedUsers)
			user.POST("/assisted/:id", assistHandler.RequestAssistedUserLink)
			user.PUT("/assisted/:id", assistHandler.AcceptAssistedUserRequest)
			user.DELETE("/assisted/:id", assistHandler.DeleteAssistedLink)

			user.GET("/assistants/requests", assistHandler.ListAssistantsRequests)
			user.GET("/assistants", assistHandler.ListAssistants)
			user.POST("/assistants/:id", assistHandler.RequestAssistantLink)
			user.PUT("/assistants/:id", assistHandler.AcceptAssistantRequest)
			user.DELETE("/assistants/:id", assistHandler.DeleteAssistantLink)
		}
		entry := v1.Group("/entries", authMiddleware())
		{
			entry.POST("/", entryHandler.Create)
			entry.GET("/", entryHandler.List)
			entry.GET("/:id", entryHandler.Find)
			entry.PUT("/:id", entryHandler.Update)
			entry.DELETE("/:id", entryHandler.Delete)
		}
		dino := v1.Group("/dinos", authMiddleware())
		{
			dino.POST("/", dinoHandler.Create)
			dino.GET("/", dinoHandler.List)
			dino.GET("/:id", dinoHandler.Find)
			dino.PUT("/:id", dinoHandler.Update)
			dino.DELETE("/:id", dinoHandler.Delete)
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
