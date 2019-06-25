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

func writer(lvl zapcore.Level, a ...interface{}) {
	var msg = fmt.Sprint(a...)
	if !doLoggerInitialized {
		InitiatLogger()
	}
	if ce := internalLogger.Check(lvl, msg); ce != nil {
		ce.Entry.Caller = zapcore.NewEntryCaller(runtime.Caller(2))
		ce.Write()
	}
}

func writerf(lvl zapcore.Level, format string, prm ...interface{}) {
	var msg = fmt.Sprintf(format, prm...)
	if !doLoggerInitialized {
		InitiatLogger()
	}
	if ce := internalLogger.Check(lvl, msg); ce != nil {
		ce.Entry.Caller = zapcore.NewEntryCaller(runtime.Caller(2))
		ce.Write()
	}
}

// Debug :
func Debug(a ...interface{}) {
	writer(zap.DebugLevel, a...)
}

// Debugf :
func Debugf(format string, prm ...interface{}) {
	writerf(zap.DebugLevel, format, prm...)
}

// Info :
func Info(a ...interface{}) {
	writer(zap.InfoLevel, a...)
}

// Infof :
func Infof(format string, prm ...interface{}) {
	writerf(zap.InfoLevel, format, prm...)
}

// Warn :
func Warn(a ...interface{}) {
	writer(zap.WarnLevel, a...)
}

// Warnf :
func Warnf(format string, prm ...interface{}) {
	writerf(zap.WarnLevel, format, prm...)
}

// Error :
func Error(a ...interface{}) {
	writer(zap.ErrorLevel, a...)
}

// Errorf :
func Errorf(format string, prm ...interface{}) {
	writerf(zap.ErrorLevel, format, prm...)
}

// Fatal :
func Fatal(a ...interface{}) {
	writer(zap.FatalLevel, a...)
}

// Fatalf :
func Fatalf(format string, prm ...interface{}) {
	writerf(zap.FatalLevel, format, prm...)
}

// Sync :
func Sync() {
	if !doLoggerInitialized {
		InitiatLogger()
	}
	internalLogger.Sync()
}
