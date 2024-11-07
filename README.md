# logi

[![Go PKG](https://raw.githubusercontent.com/rakunlabs/.github/main/assets/badges/gopkg.svg)](https://pkg.go.dev/github.com/rakunlabs/logi)

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

Context logging

```go
ctx := logi.WithContext(context.Background(), slog.With(slog.String("component", "example")))

logi.Ctx(ctx).Info("This is a log message")
```
