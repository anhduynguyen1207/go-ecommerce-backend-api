package logger

import (
	"os"

	"github.com/anhduynguyen1207/go-ecommerce-backend-api/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.Log_level
	// debug -> info -> warn -> error -> fatal -> panic
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.File_log_name,
		MaxSize:    config.Max_age, // megabytes
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age,  //days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)
	// logger := zap.New(core, zap.AddCaller())
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

// hàm custome lại fomart log theo mong muốn
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// 1737197083.6017373 => 2021-09-17T15:38:03.6017373+07:00
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//ts => time
	encodeConfig.TimeKey = "time"
	// from info => INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//cli/main.log.go:24 => file:line
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
