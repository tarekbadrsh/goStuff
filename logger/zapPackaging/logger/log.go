package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("|%s|", t.Format("2006-01-02T15:04:05")))
}

func filelogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%s", t.Format("2006-01-02T15:04:05")))
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%v]", level.CapitalString()))
}

var internalLogger *zap.Logger
var doLoggerInitialized bool

// InitiatLogger :
func InitiatLogger() {
	terminalEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:     "M",
		LevelKey:       "L",
		TimeKey:        "T",
		NameKey:        "N",
		CallerKey:      "C",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    customLevelEncoder,
		EncodeTime:     syslogTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	terminalOutput := zapcore.AddSync(os.Stderr)

	fileEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:     "M",
		LevelKey:       "L",
		TimeKey:        "T",
		NameKey:        "N",
		CallerKey:      "C",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     filelogTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	fileOutput := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs.json",
		MaxSize:    10, // megabytes
		MaxBackups: 100,
		MaxAge:     28, // days
		Compress:   true,
	})

	InfoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	ErrorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(terminalEncoder, terminalOutput, InfoLevel),
		zapcore.NewCore(fileEncoder, fileOutput, ErrorLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	logger.Info("logger.info")
	logger.Error("logger.Error ")
	internalLogger = logger
	doLoggerInitialized = true
}

// Debug :
func Debug(msg string) {
	if !doLoggerInitialized {
		InitiatLogger()
	}
	internalLogger.Debug(msg)
}

// Info :
func Info(msg string) {
	if !doLoggerInitialized {
		InitiatLogger()
	}
	internalLogger.Info(msg)
}

// Warn :
func Warn(msg string) {
	if !doLoggerInitialized {
		InitiatLogger()
	}
	internalLogger.Warn(msg)
}

// Error :
func Error(msg string) {
	if !doLoggerInitialized {
		InitiatLogger()
	}
	internalLogger.Error(msg)
}

// Fatal :
func Fatal(msg string) {
	if !doLoggerInitialized {
		InitiatLogger()
	}
	internalLogger.Fatal(msg)
}

// Sync :
func Sync() {
	if !doLoggerInitialized {
		InitiatLogger()
	}
	internalLogger.Sync()
}
