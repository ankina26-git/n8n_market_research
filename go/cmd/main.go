package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	//各アプリルーティング
	analytics "n8n_project_go/app/analytics/v1"
	searchconsole "n8n_project_go/app/searchconsole/v1"
	user "n8n_project_go/app/user/v1"
	"n8n_project_go/config"
	"n8n_project_go/logger"

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

	//ログ初期化
	logger.Init("debug")

	//ルーティンググループ
	//ユーザールーティング
	app.Use(logger.Middleware("user"))
	user.RegisterRoutes(app.Group("/api/user"))

	//アナリティクスルーティング
	app.Use(logger.Middleware("analytics"))
	analytics.RegisterRoutes(app.Group("/api/analytics"))

	//サーチコンソールルーティング
	app.Use(logger.Middleware("searchconsole"))
	searchconsole.RegisterRoutes(app.Group("/api/searchconsole"))

	//接続結果ログ
	log.Printf("🚀 Starting Fiber on port %s", port)
	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("❌ Failed to start server: %v", err)
		}
	}()

	defer func() {
		for name, l := range logger.AppLoggers {
			if err := l.Sync(); err != nil {
				log.Printf("⚠️ failed to sync logger for %s: %v", name, err)
			}
		}
	}()

	//失敗時シャットダウン
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Println("🛑 Gracefully shutting down...")
	_ = app.Shutdown()
}
