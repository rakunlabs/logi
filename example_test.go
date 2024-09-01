package logi_test

import (
	"log/slog"
	"os"

	"github.com/rakunlabs/logi"
)

func ExampleInitializeLog() {
	logi.InitializeLog(logi.WithTimeStamp("-"), logi.WithCaller(false), logi.WithWriter(os.Stdout))

	_ = logi.SetLogLevel("ERROR")

	slog.Error("This is an error message")
	slog.Info("Yet another log message")

	// Output:
	// {"time":"-","level":"ERROR","msg":"This is an error message"}
}
