package logi

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

var (
	EnvPretty        = "LOG_PRETTY"
	EnvLevel         = "LOG_LEVEL"
	TimeFormat       = time.RFC3339Nano
	TimePrettyFormat = "2006-01-02 15:04:05 MST"
)

// InitializeLog choice between json format or common format.
// LOG_PRETTY boolean environment value always override the decision.
// Override with some option argument.
func InitializeLog(opts ...Option) {
	logger := Logger(opts...)

	slog.SetDefault(logger)
}

func Logger(opts ...Option) *slog.Logger {
	opt := &option{
		Level:  slog.LevelInfo.String(),
		Caller: true,
		Writer: os.Stderr,
		Pretty: SelectAuto,
	}
	opt.apply(opts...)

	// ///////////////////////////////////
	pretty := isPretty(opt.Pretty, opt.Writer)

	// ///////////////////////////////////
	levelStr := checkLevel(opt.Level)

	var sloglevel slog.Level
	_ = sloglevel.UnmarshalText([]byte(levelStr))

	// ///////////////////////////////////
	var logger *slog.Logger

	if pretty {
		logger = slog.New(
			tint.NewHandler(
				opt.Writer,
				&tint.Options{
					AddSource:  opt.Caller,
					Level:      sloglevel,
					TimeFormat: timeFormat(opt.TimeFormat, pretty),
				},
			),
		)
	} else {
		logger = slog.New(
			slog.NewJSONHandler(
				opt.Writer,
				&slog.HandlerOptions{
					AddSource: opt.Caller,
					Level:     sloglevel,
				},
			),
		)
	}

	return logger
}
