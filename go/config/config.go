package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		AppName:      "Fiber Production App", //
		Prefork:      true,                   // マルチプロセスモード（CPUコア数に応じたプロセス分岐）
		ReadTimeout:  10 * time.Second,       // リクエストの読み取りタイムアウト 10秒
		WriteTimeout: 10 * time.Second,       // レスポンスの書き込みタイムアウト 10秒
	}
}
