package logi

import (
	"io"
)

type option struct {
	Pretty     selection
	TimeFormat *string
	Caller     bool
	Level      string
	Writer     io.Writer
}

func (o *option) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(opt *option)

func WithTimeStamp(timeFormat string) Option {
	return func(opt *option) {
		opt.TimeFormat = &timeFormat
	}
}

func WithCaller(caller bool) Option {
	return func(opt *option) {
		opt.Caller = caller
	}
}

func WithPretty(pretty selection) Option {
	return func(opt *option) {
		opt.Pretty = pretty
	}
}

// WithLevel sets the log level.
//
// The level must be one of the following:
// "DEBUG", "INFO", "WARN", "ERROR".
func WithLevel(level string) Option {
	return func(options *option) {
		options.Level = level
	}
}

func WithWriter(w io.Writer) Option {
	return func(options *option) {
		options.Writer = w
	}
}
