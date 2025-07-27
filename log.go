package logi

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"strings"
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
// Set the default logger and return it.
//   - If the LOG_PRETTY environment variable is set to true, the pretty format will be used.
//   - If the LOG_LEVEL environment variable is set, the log level will be set.
func InitializeLog(opts ...Option) *slog.Logger {
	logger := Logger(opts...)

	slog.SetDefault(logger)

	return logger
}

type handlerWrapper struct {
	Level *slog.Level
	slog.Handler
}

func (h *handlerWrapper) SetLogLevel(levelStr string) error {
	return h.Level.UnmarshalText([]byte(levelStr))
}

type setLeveler interface {
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
	levelStr := strings.ToUpper(checkLevel(opt.Level))

	sloglevel := new(slog.Level)
	_ = sloglevel.UnmarshalText([]byte(levelStr))

	// ///////////////////////////////////
	var logger *slog.Logger

	tFormat := timeFormat(opt.TimeFormat, pretty)

	if pretty {
		logger = slog.New(
			&handlerWrapper{
				Level: sloglevel,
				Handler: tint.NewHandler(
					opt.Writer,
					&tint.Options{
						AddSource:  opt.Caller,
						Level:      sloglevel,
						TimeFormat: tFormat,
					},
				),
			},
		)
	} else {
		logger = slog.New(
			&handlerWrapper{
				Level: sloglevel,
				Handler: slog.NewJSONHandler(
					opt.Writer,
					&slog.HandlerOptions{
						AddSource: opt.Caller,
						Level:     sloglevel,
						ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
							if a.Key == slog.TimeKey {
								a.Value = slog.StringValue(a.Value.Time().Format(tFormat))
							}

							return a
						},
					},
				),
			},
		)
	}

	return logger
}

// SetLogLevel set the log level of the default logger.
//   - Just work if the handler implements `SetLogLevel(levelStr string) error` function.
func SetLogLevel(levelStr string) error {
	if wrapper, ok := slog.Default().Handler().(setLeveler); ok {
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
