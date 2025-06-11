package config

import "github.com/gofiber/fiber/v2"

func FiberConfig() fiber.Config {
	return fiber.Config{
		AppName:      "Fiber Production App",
		Prefork:      true,                    // マルチコア対応（リバースプロキシ時）
		ReadTimeout:  10 * 1000 * 1000 * 1000, // 10秒
		WriteTimeout: 10 * 1000 * 1000 * 1000,
	}
}
