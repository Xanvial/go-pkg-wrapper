### Sample usage

```
// using default logger
log.Info(context.Background(), "test info log", errors.New("test error"), log.KV{
    "data1": 123,
    "data2": "data 2 content",
})
log.Printf("sample data, value:%d", 123)

// override default logger
log.SetLogger(log.NewZerolog(log.Config{
    LogLevel:       log.LogLevelTrace,
    TimeFormat:     log.LogTimeFormatTimestamp,
    OutputType:     log.OutputFile,
    OutputFilePath: "test.log",
}))
log.Info(context.Background(), "test debug log123123", errors.New("test error"), log.KV{
    "data3": 111,
    "data4": "data 4 content",
})
```