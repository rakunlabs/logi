# logi

Log initializer for golang's slog.  
If terminal is detected, it will use colorized output else it will use JSON output.

```sh
go get github.com/rakunlabs/logi
```

## Usage

```go
logi.InitializeLog()

slog.Error("This is an error message")
slog.Info("Yet another log message")
```

For setting global log level, uses level parse from slog package.

```go
logi.SetLogLevel("ERROR")
```
