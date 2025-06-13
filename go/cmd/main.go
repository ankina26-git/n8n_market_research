package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	//各アプリルーティング
	analytics "n8n_project_go/app/Analytics/v1"
	user "n8n_project_go/app/User/v1"
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
	//環境変数デフォルト設定
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	//fiber初期化
	app := fiber.New(FiberConfig())

	//DB接続
	config.PostgresDB()
	config.DB.AutoMigrate(&user.User{})

	//ルーティンググループ
	user.RegisterRoutes(app.Group("/api/user"))
	analytics.RegisterRoutes(app.Group("/api/analytics"))

	//接続結果ログ
	log.Printf("🚀 Starting Fiber on port %s", port)
	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("❌ Failed to start server: %v", err)
		}
	}()

	//失敗時シャットダウン
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Println("🛑 Gracefully shutting down...")
	_ = app.Shutdown()
}
