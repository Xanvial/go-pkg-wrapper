package log

import "context"

type Log interface {
	Print(msg string)
	Printf(msg string, v ...any)

	Trace(ctx context.Context, msg string, err error, fields KV)
	Info(ctx context.Context, msg string, err error, fields KV)
	Warn(ctx context.Context, msg string, err error, fields KV)
	Error(ctx context.Context, msg string, err error, fields KV)
	Debug(ctx context.Context, msg string, err error, fields KV)
	Panic(ctx context.Context, msg string, err error, fields KV)
	Fatal(ctx context.Context, msg string, err error, fields KV)
}
