package logi_test

import (
	"log/slog"

	"github.com/rakunlabs/logi"
)

func ExampleInitializeLog() {
	logi.InitializeLog()

	_ = logi.SetLogLevel("ERROR")

	slog.Error("This is an error message")
	slog.Info("Yet another log message")
	// {"time":"2024-05-08T23:57:36.393939445+02:00","level":"INFO","source":{"function":"github.com/rakunlabs/logi_test.ExampleLogi","file":"/logi/example_test.go","line":13},"msg":"Yet another log message"}
}
