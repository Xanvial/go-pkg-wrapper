package log

import (
	"context"
	"io"
	stdLog "log"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type zerologger struct {
	logger zerolog.Logger
}

func NewZerolog(cfg Config) Log {

	var (
		targetOutput io.Writer
		err          error
	)

	// set output
	switch cfg.OutputType {
	case OutputStdout:
		targetOutput = os.Stdout
	case OutputStderr:
		targetOutput = os.Stderr
	case OutputFile:
		targetOutput, err = os.OpenFile(
			cfg.OutputFilePath,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		if err != nil {
			stdLog.Println("failed to create log file, output to stdout instead. err:", err)
			targetOutput = os.Stdout
		}
	}

	builder := zerolog.New(targetOutput).With()

	switch cfg.TimeFormat {
	case LogTimeFormatTimestamp:
		zerolog.TimeFieldFormat = time.RFC3339
		builder = builder.Timestamp()
	case LogTimeFormatUnix:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		builder = builder.Timestamp()
	case LogTimeFormatDisable:
		// do nothing
	}

	var zeroLogLevel zerolog.Level
	switch cfg.LogLevel {
	case LogLevelTrace:
		zeroLogLevel = zerolog.TraceLevel
	case LogLevelDebug:
		zeroLogLevel = zerolog.DebugLevel
	case LogLevelInfo:
		zeroLogLevel = zerolog.InfoLevel
	case LogLevelWarn:
		zeroLogLevel = zerolog.WarnLevel
	case LogLevelError:
		zeroLogLevel = zerolog.ErrorLevel
	case LogLevelPanic:
		zeroLogLevel = zerolog.PanicLevel
	case LogLevelFatal:
		zeroLogLevel = zerolog.FatalLevel
	default:
		// use info level as default
		zeroLogLevel = zerolog.InfoLevel
	}

	return &zerologger{
		logger: builder.Logger().Level(zeroLogLevel),
	}
}

func (zl *zerologger) Print(msg string) {
	zl.logger.Info().Msg(msg)
}

func (zl *zerologger) Printf(msg string, v ...any) {
	zl.logger.Info().Msgf(msg, v...)
}

func (zl *zerologger) Trace(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Trace().
		Err(err).
		Fields(map[string]any(fields)).
		Msg(msg)
}

func (zl *zerologger) Info(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Info().
		Err(err).
		Fields(map[string]any(fields)).
		Msg(msg)
}

func (zl *zerologger) Warn(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Warn().
		Err(err).
		Fields(map[string]any(fields)).
		Msg(msg)
}

func (zl *zerologger) Error(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Error().
		Err(err).
		Fields(map[string]any(fields)).
		Msg(msg)
}

func (zl *zerologger) Debug(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Debug().
		Err(err).
		Fields(map[string]any(fields)).
		Msg(msg)
}

func (zl *zerologger) Panic(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Panic().
		Err(err).
		Fields(map[string]any(fields)).
		Msg(msg)
}

func (zl *zerologger) Fatal(ctx context.Context, msg string, err error, fields KV) {
	zl.logger.Fatal().
		Err(err).
		Fields(map[string]any(fields)).
		Msg(msg)
}
