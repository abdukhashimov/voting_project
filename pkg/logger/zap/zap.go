// Package zap handles creating zap logger
package zap

import (
	"time"

	"github.com/abdukhashimov/go_api/pkg/logger"
	"github.com/abdukhashimov/go_api/pkg/logger/options"
	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"
)

// RegisterLog creates a new Zap logger
func RegisterLog(cfg *options.Logging) (logger.Logger, error) {
	encoderCfg := loadEncoderConfig(cfg.DevMode, cfg.DateTimeFormat)

	globalLogLevel, err := zapcore.ParseLevel(cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	loggerConfig := zap.Config{
		Level:         zap.NewAtomicLevelAt(globalLogLevel),
		Development:   cfg.DevMode,
		Encoding:      cfg.Encoding,
		EncoderConfig: encoderCfg,
		OutputPaths:   []string{"stdout"},
		DisableCaller: !cfg.EnableCaller,
	}

	zapLog, err := loggerConfig.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	sugar := zapLog.Named(cfg.ProjectName).Sugar()

	return sugar, nil
}

// loadEncoderConfig creates zap encode options
func loadEncoderConfig(devMode bool, timeFormat string) zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()

	if devMode {
		cfg = zap.NewDevelopmentEncoderConfig()
	}

	cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(timeFormat))
	}

	return cfg
}
