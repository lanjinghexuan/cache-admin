package pkg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func Init(mode string) {
	var cfg zap.Config
	if mode == "prod" {
		cfg = zap.NewProductionConfig()
		cfg.OutputPaths = []string{"stdout", "./logs/app.log"}
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.OutputPaths = []string{"stdout"}
	}

	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zapLogger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	logger = zapLogger.Sugar()
}

// 基础方法
func Info(msg string, fields ...interface{}) {
	logger.Infof(msg, fields...)
}

func Debug(msg string, fields ...interface{}) {
	logger.Debugf(msg, fields...)
}

func Error(msg string, fields ...interface{}) {
	logger.Errorf(msg, fields...)
}

func Fatal(msg string, fields ...interface{}) {
	logger.Fatalf(msg, fields...)
}

// 返回原始 zap.Logger（可用于更复杂场景）
func Raw() *zap.SugaredLogger {
	return logger
}
