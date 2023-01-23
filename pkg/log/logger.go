package log

import (
	"context"
)

// Default logger, can be overridden when needed
var logger = NewZerolog(Config{
	OutputType: OutputStdout,
	LogLevel:   LogLevelInfo,
	TimeFormat: LogTimeFormatUnix,
})

func SetLogger(log Log) {
	logger = log
}

func Print(msg string) {
	logger.Print(msg)
}

func Printf(msg string, v ...any) {
	logger.Printf(msg, v...)
}

func Trace(ctx context.Context, msg string, err error, fields KV) {
	logger.Trace(ctx, msg, err, fields)
}

func Info(ctx context.Context, msg string, err error, fields KV) {
	logger.Info(ctx, msg, err, fields)
}

func Warn(ctx context.Context, msg string, err error, fields KV) {
	logger.Warn(ctx, msg, err, fields)
}

func Error(ctx context.Context, msg string, err error, fields KV) {
	logger.Error(ctx, msg, err, fields)
}

func Debug(ctx context.Context, msg string, err error, fields KV) {
	logger.Debug(ctx, msg, err, fields)
}

func Panic(ctx context.Context, msg string, err error, fields KV) {
	logger.Panic(ctx, msg, err, fields)
}

func Fatal(ctx context.Context, msg string, err error, fields KV) {
	logger.Fatal(ctx, msg, err, fields)
}
