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

const TOKEN_DURATION = 72

func main() {
	appName := os.Getenv("APP_NAME")
	env := os.Getenv("APP_ENV")
	dbConn := os.Getenv("DB_CONNECTION")
	tokenKey := os.Getenv("TOKEN_KEY")
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

	tokenProvider := token.NewTokenProvider(TOKEN_DURATION, tokenKey)

	// Dependency injection
	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService, tokenProvider)

	// Auth
	authService := service.NewAuthService(userRepo, tokenProvider)
	authHandler := handlers.NewAuthHandler(authService) //TODO: This should receive tokenProvider, not the service

	// Assist
	assistRepo := repository.NewAssistRepository(db)
	assistService := service.NewAssistService(assistRepo)
	assistHandler := handlers.NewAssistHandler(assistService, tokenProvider)

	// Entry
	entryRepo := repository.NewEntryRepository(db)
	entryService := service.NewEntryService(entryRepo)
	entryHandler := handlers.NewEntryHandler(entryService, tokenProvider)

	router, err := handlers.NewRouter(
		tokenProvider.Authenticate,
		*userHandler,
		*assistHandler,
		*authHandler,
		*entryHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
	}

	// Start server
	slog.Info("Starting HTTP server", "listen_address", listenAddr)
	if env == "PROD" {
		router.ServeLambda()
	} else {
		router.Serve(listenAddr)
	}
}
