package logger

import (
	"fmt"
	"os"
	"runtime"
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
	internalLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	doLoggerInitialized = true
}

// Debug :
func Debug(format string, prm ...interface{}) {
	if !doLoggerInitialized {
		InitiatLogger()
	}
	var msg = fmt.Sprintf(format, prm...)
	if ce := internalLogger.Check(zap.DebugLevel, msg); ce != nil {
		ce.Entry.Caller = zapcore.NewEntryCaller(runtime.Caller(1))
		ce.Write()
	}
}

// Info :
func Info(format string, prm ...interface{}) {
	if !doLoggerInitialized {
		InitiatLogger()
	}
	var msg = fmt.Sprintf(format, prm...)
	if ce := internalLogger.Check(zap.InfoLevel, msg); ce != nil {
		ce.Entry.Caller = zapcore.NewEntryCaller(runtime.Caller(1))
		ce.Write()
	}
}

// Warn :
func Warn(format string, prm ...interface{}) {
	if !doLoggerInitialized {
		InitiatLogger()
	}
	var msg = fmt.Sprintf(format, prm...)
	if ce := internalLogger.Check(zap.WarnLevel, msg); ce != nil {
		ce.Entry.Caller = zapcore.NewEntryCaller(runtime.Caller(1))
		ce.Write()
	}
}

// Error :
func Error(format string, prm ...interface{}) {
	if !doLoggerInitialized {
		InitiatLogger()
	}
	var msg = fmt.Sprintf(format, prm...)
	if ce := internalLogger.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Entry.Caller = zapcore.NewEntryCaller(runtime.Caller(1))
		ce.Write()
	}
}

// Fatal :
func Fatal(format string, prm ...interface{}) {
	if !doLoggerInitialized {
		InitiatLogger()
	}
	var msg = fmt.Sprintf(format, prm...)
	if ce := internalLogger.Check(zap.FatalLevel, msg); ce != nil {
		ce.Entry.Caller = zapcore.NewEntryCaller(runtime.Caller(1))
		ce.Write()
	}
}

// Sync :
func Sync() {
	if !doLoggerInitialized {
		InitiatLogger()
	}
	internalLogger.Sync()
}
