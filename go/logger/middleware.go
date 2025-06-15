package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var AppLoggers = make(map[string]*zap.Logger)

func AppLogger(appName string) *zap.Logger {
	if logger, ok := AppLoggers[appName]; ok {
		return logger
	}

	logDir := filepath.Join("app", appName, "logs")
	_ = os.MkdirAll(logDir, os.ModePerm)

	logFile := filepath.Join(logDir, fmt.Sprintf("%s.log", appName))

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{logFile, "stdout"}
	cfg.ErrorOutputPaths = []string{logFile, "stderr"}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	AppLoggers[appName] = logger
	return logger
}

// Fiber用の共通ロガーミドルウェア
func Middleware(appName string) fiber.Handler {
	logger := AppLogger(appName) // ← 各アプリ専用ロガーを生成

	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		logger.Info("API Request",
			zap.String("method", c.Method()),
			zap.String("path", c.OriginalURL()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
			zap.String("ip", c.IP()),
			zap.String("user_agent", c.Get("User-Agent")),
		)

		return err
	}
}
