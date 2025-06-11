package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	user_v1 "n8n_project_go/app/user/v1"
	"n8n_project_go/config"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		AppName: "Fiber Production App",
		Prefork: false, // Dockerで安定稼働
	}
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New(FiberConfig())

	config.PostgresDB()
	config.DB.AutoMigrate(&user_v1.User{})
	user_v1.RegisterRoutes(app.Group("/api/v1/user"))

	log.Printf("🚀 Starting Fiber on port %s", port)
	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("❌ Failed to start server: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Println("🛑 Gracefully shutting down...")
	_ = app.Shutdown()
}
