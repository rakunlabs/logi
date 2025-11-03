package logi_test

import (
	"context"
	"log/slog"
	"os"

	"github.com/rakunlabs/logi"
)

func ExampleInitializeLog() {
	// stdout writer for output test
	logi.InitializeLog(logi.WithTimeStamp("-"), logi.WithCaller(false), logi.WithWriter(os.Stdout))

	_ = logi.SetLogLevel("ERROR")

	slog.Error("This is an error message")
	slog.Info("Yet another log message")

	// Output:
	// {"time":"-","level":"ERROR","msg":"This is an error message"}
}

func ExampleWithContext() {
	// stdout writer for output test
	logi.InitializeLog(logi.WithTimeStamp("-"), logi.WithCaller(false), logi.WithWriter(os.Stdout))

	ctx := logi.WithContext(context.Background(), slog.With(slog.String("component", "example")))

	logi.Ctx(ctx).Info("This is a log message", "object", `{"test": 1234}`, "address", "[::]:8080")

	logi.Ctx(context.Background()).Info("Empty context")

	// Output:
	// {"time":"-","level":"INFO","msg":"This is a log message","component":"example","object":{"test":1234},"address":"[::]:8080"}
	// {"time":"-","level":"INFO","msg":"Empty context"}
}

func Example() {
	// stdout writer for output test
	logi.InitializeLog(logi.WithTimeStamp("-"), logi.WithCaller(false), logi.WithWriter(os.Stdout), logi.WithPrettyStr("true"))

	slog.Info("This is a log message", "object", `{"test": 1234, "inner": {"key": "value"}}`, "address", "[::]:8080")

	// Output:
	// [2m-[0m [92mINF[0m This is a log message [2mobject=[0m{"test": 1234, "inner": {"key": "value"}} [2maddress=[0m[::]:8080
}
