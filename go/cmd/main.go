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

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	// Fiberインスタンス作成（Preforkモード有効でもOK）
	app := fiber.New(config.FiberConfig())

	// Prefork時はこのif文で "Master" プロセスだけが初期化処理を行う
	if fiber.IsChild() {
		log.Println("👶 Child process detected → skipping DB init and route setup")
	} else {
		// マスタープロセスでのみ初期化
		config.PostgresDB()
		user_v1.RegisterRoutes(app.Group("/api/v1/user"))
	}

	// 1回だけListen（非同期でなくてもOK）
	log.Printf("🚀 Starting server on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}

	// グレースフルシャットダウン（Fiberが内包している場合 Prefork では動作保証されないこともある）
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Println("🛑 Shutting down gracefully...")
	_ = app.Shutdown()
}
