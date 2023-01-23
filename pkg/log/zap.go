package log

import (
	"context"
	"fmt"
	stdLog "log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zaplog struct {
	logger zap.Logger
}

func NewZaplog(cfg Config) Log {

	var zapLogLevel zapcore.Level
	switch cfg.LogLevel {
	case LogLevelTrace:
		zapLogLevel = zapcore.DebugLevel
	case LogLevelDebug:
		zapLogLevel = zapcore.DebugLevel
	case LogLevelInfo:
		zapLogLevel = zapcore.InfoLevel
	case LogLevelWarn:
		zapLogLevel = zapcore.WarnLevel
	case LogLevelError:
		zapLogLevel = zapcore.ErrorLevel
	case LogLevelPanic:
		zapLogLevel = zapcore.PanicLevel
	case LogLevelFatal:
		zapLogLevel = zapcore.FatalLevel
	default:
		// use info level as default
		zapLogLevel = zapcore.InfoLevel
	}

	// set output
	var targetOutput string
	switch cfg.OutputType {
	case OutputStdout:
		targetOutput = "stdout"
	case OutputStderr:
		targetOutput = "stderr"
	case OutputFile:
		targetOutput = cfg.OutputFilePath
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time" // parity with zerolog format

	switch cfg.TimeFormat {
	case LogTimeFormatTimestamp:
		encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	case LogTimeFormatUnix:
		encoderCfg.EncodeTime = zapcore.EpochTimeEncoder
	case LogTimeFormatDisable:
		encoderCfg.TimeKey = ""
	}

	zapConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLogLevel),
		Development:       false,
		DisableStacktrace: true,
		DisableCaller:     true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    encoderCfg,
		OutputPaths:      []string{targetOutput},
		ErrorOutputPaths: []string{targetOutput},
	}
	logger, err := zapConfig.Build()
	if err != nil {
		stdLog.Println("failed to initialize log with provided config, using default configuration instead. err:", err)
		zapProductionLog, _ := zap.NewProductionConfig().Build()
		return &zaplog{logger: *zapProductionLog}
	}

	return &zaplog{
		logger: *logger,
	}
}

func (zl *zaplog) Print(msg string) {
	zl.logger.Info(msg)
}

func (zl *zaplog) Printf(msg string, v ...any) {
	zl.logger.Info(fmt.Sprintf(msg, v...))
}

func (zl *zaplog) Trace(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Debug(msg, mergeZapFields(fields, err)...)
}

func (zl *zaplog) Info(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Info(msg, mergeZapFields(fields, err)...)
}

func (zl *zaplog) Warn(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Warn(msg, mergeZapFields(fields, err)...)
}

func (zl *zaplog) Error(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Error(msg, mergeZapFields(fields, err)...)
}

func (zl *zaplog) Debug(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Debug(msg, mergeZapFields(fields, err)...)
}

func (zl *zaplog) Panic(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Panic(msg, mergeZapFields(fields, err)...)
}

func (zl *zaplog) Fatal(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Fatal(msg, mergeZapFields(fields, err)...)
}

func mergeZapFields(m KV, err error) []zap.Field {
	fields := make([]zap.Field, 0, len(m))
	for k, v := range m {
		fields = append(fields, zap.Any(k, v))
	}

	if err != nil {
		fields = append(fields, zap.String("error", err.Error()))
	}

	return fields
}
