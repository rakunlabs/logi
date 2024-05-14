package logi

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/lmittmann/tint"
)

var (
	EnvPretty        = "LOG_PRETTY"
	EnvLevel         = "LOG_LEVEL"
	TimeFormat       = time.RFC3339Nano
	TimePrettyFormat = "2006-01-02 15:04:05 MST"
	ErrorKey         = "error"
)

// InitializeLog choice between json format or common format.
// LOG_PRETTY boolean environment value always override the decision.
// Override with some option argument.
func InitializeLog(opts ...Option) {
	logger := Logger(opts...)

	slog.SetDefault(logger)
}

type HandlerWrapper struct {
	Level *slog.Level
	slog.Handler
}

func (h *HandlerWrapper) SetLogLevel(levelStr string) error {
	return h.Level.UnmarshalText([]byte(levelStr))
}

type SetLeveler interface {
	SetLogLevel(levelStr string) error
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

	sloglevel := new(slog.Level)
	_ = sloglevel.UnmarshalText([]byte(levelStr))

	// ///////////////////////////////////
	var logger *slog.Logger

	if pretty {
		logger = slog.New(
			&HandlerWrapper{
				Level: sloglevel,
				Handler: tint.NewHandler(
					opt.Writer,
					&tint.Options{
						AddSource:  opt.Caller,
						Level:      sloglevel,
						TimeFormat: timeFormat(opt.TimeFormat, pretty),
					},
				),
			},
		)
	} else {
		logger = slog.New(
			&HandlerWrapper{
				Level: sloglevel,
				Handler: slog.NewJSONHandler(
					opt.Writer,
					&slog.HandlerOptions{
						AddSource: opt.Caller,
						Level:     sloglevel,
					},
				),
			},
		)
	}

	return logger
}

// SetLogLevel set the log level of the default logger.
//   - Just work if the handler implements SetLeveler interface.
func SetLogLevel(levelStr string) error {
	if wrapper, ok := slog.Default().Handler().(SetLeveler); ok {
		if err := wrapper.SetLogLevel(levelStr); err != nil {
			return err
		}
	}

	return nil
}

// Log for without level check.
func Log(msg string, args ...any) {
	var pc uintptr
	if true {
		var pcs [1]uintptr
		// skip [runtime.Callers, this function, this function's caller]
		runtime.Callers(2, pcs[:])
		pc = pcs[0]
	}

	r := slog.NewRecord(time.Now(), slog.LevelInfo, msg, pc)
	r.Add(args...)

	_ = slog.Default().Handler().Handle(context.Background(), r)
}
