package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d ", "anhduy", 20)

	// //logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "anhduy"), zap.Int("age", 20))

	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")

	// //Development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")

	// //Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	encoder := getEncoderLog()
	sync := getWriteSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
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

func getWriteSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_RDONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
