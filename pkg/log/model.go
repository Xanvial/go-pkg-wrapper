package log

type KV map[string]interface{}

type LogLevel int

const (
	LogLevelUnknown = iota
	LogLevelTrace
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelPanic
	LogLevelFatal
)

type LogTimeFormat int

const (
	LogTimeFormatDisable = iota
	LogTimeFormatStamp
)

type Output int

const (
	OutputStdout = iota
	OutputStderr
	OutputFile
)

type Config struct {
	LogLevel       LogLevel
	TimeFormat     LogTimeFormat
	OutputType     Output
	OutputFilePath string // only if output type is set to OutputFile
}
