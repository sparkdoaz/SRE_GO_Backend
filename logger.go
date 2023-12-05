package main

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func getEncoder() zapcore.Encoder {
	// return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig) // 使用 ConsoleEncoder
}

func getLogWriter() zapcore.WriteSyncer {
	// 使用 MultiWriteSyncer 同時輸出到檔案和控制台
	fileWriter := getLogFileWriter()
	consoleWriter := zapcore.Lock(os.Stdout)
	return zapcore.NewMultiWriteSyncer(fileWriter, consoleWriter)
}

func getLogFileWriter() zapcore.WriteSyncer {
	// 根據時間生成檔案名稱
	currentTime := time.Now()
	logFileName := currentTime.Format("2006-01-02_15-04-05") + ".log"
	file, _ := os.Create("./log/" + logFileName)
	return zapcore.AddSync(file)
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}
