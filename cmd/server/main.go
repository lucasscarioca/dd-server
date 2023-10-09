package main

import (
	"log/slog"
	"os"

	"github.com/lucasscarioca/dinodiary/internal/adapters/handlers"
	repository "github.com/lucasscarioca/dinodiary/internal/adapters/repository"
	"github.com/lucasscarioca/dinodiary/internal/adapters/token"
	"github.com/lucasscarioca/dinodiary/internal/core/service"
)

func init() {
	var logHandler *slog.JSONHandler

	env := os.Getenv("APP_ENV")
	if env == "PROD" {
		logHandler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}

func main() {
	appName := os.Getenv("APP_NAME")
	env := os.Getenv("APP_ENV")
	dbConn := os.Getenv("DB_CONNECTION")
	tokenSecret := os.Getenv("SECRET")
	listenAddr := ":" + os.Getenv("PORT")

	slog.Info("Starting the application", "app", appName, "env", env)

	// Init DB
	db, err := repository.NewDB()
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database", "db", dbConn)

	//TODO: Init Cache service

	tokenProvider := token.NewTokenProvider(72, tokenSecret)

	// Dependency injection
	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Auth
	authService := service.NewAuthService(userRepo, tokenProvider)
	authHandler := handlers.NewAuthHandler(authService)

	router, err := handlers.NewRouter(
		tokenProvider,
		*userHandler,
		*authHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
	}

	// Start server
	slog.Info("Starting HTTP server", "listen_address", listenAddr)
	router.Serve(listenAddr)
}
