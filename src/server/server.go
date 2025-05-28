package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/one-d-plate/go-boilerplate.git/src/bootstrap"
	"github.com/one-d-plate/go-boilerplate.git/src/config"
	"github.com/one-d-plate/go-boilerplate.git/src/pkg"
	"github.com/one-d-plate/go-boilerplate.git/src/route"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) Run() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "auth App v1.0.1",
	})

	app.Use(cors.New(*config.CorsConfig()))

	// Hanya proses master (parent) yang menangani sinyal
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	if !fiber.IsChild() {
		pkg.WaitForSignalAndShutdown(app)
	}

	db := bootstrap.NewDatabase()
	dbConnect, err := db.Connect()
	if err != nil {
		pkg.Logger.Error("Database connection failed ", err)
		return
	}

	redis := bootstrap.NewRedisClient()
	redisClient, err := redis.Connect()
	if err != nil {
		return
	}

	err = route.RouteRegistry(app, dbConnect, redisClient)
	if err != nil {
		return
	}

	if err := app.Listen(":8080"); err != nil {
		pkg.Logger.Fatalf("Failed to start server: %v", err)
	}
}
