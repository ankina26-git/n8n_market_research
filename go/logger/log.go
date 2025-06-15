package logger

import (
	"go.uber.org/zap"
)

var baseLogger *zap.Logger

// 初期化（mainで呼び出す）
func Init(level string) {
	var cfg zap.Config
	if level == "debug" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}
	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	baseLogger = l
}

// ベースロガー
func Base() *zap.Logger {
	return baseLogger
}

// Appごとのサブロガー（アプリ識別子付き ログ出力の詳細はmiddlewareに）
func WithApp(appName string) *zap.Logger {
	return baseLogger.With(zap.String("app", appName))
}
