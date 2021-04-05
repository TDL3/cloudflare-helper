package log

import (
	"github.com/tdl3/cloudflare-helper/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init(cfg models.Config) {
	zapCfg := zap.NewProductionConfig()
	level := handleLevel(cfg)
	zapCfg.Level.SetLevel(level)
	if level != zapcore.DebugLevel {
		zapCfg.DisableCaller = true
	}
	encoding := handleEncoding(cfg)
	zapCfg.Encoding = encoding
	zapCfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006/01/02 15:04:05")
	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := zapCfg.Build()
	if err != nil {
		panic(err)
	}
	// defer logger.Sync()
	zap.ReplaceGlobals(logger)
	// defer undo()

	zap.L().Debug("replaced zap's global loggers")
}

func handleLevel(cfg models.Config) (level zapcore.Level) {
	switch cfg.Log.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func handleEncoding(cfg models.Config) (encoding string) {
	encodings := []string{"console", "json"}
	for _, encoding := range encodings {
		if encoding == cfg.Log.Encoding {
			return encoding
		}
	}
	return "console"
}
