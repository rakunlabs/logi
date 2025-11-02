package logi

import (
	"io"
)

type option struct {
	Pretty          selection
	TimeFormat      *string
	Caller          bool
	Level           string
	Writer          io.Writer
	ParseJSONString bool
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

func WithPrettyStr(pretty string) Option {
	return func(opt *option) {
		opt.Pretty = prettySelection(pretty)
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

// WithParseJSONString enables or disables automatic parsing of JSON strings.
// When enabled, string values that look like JSON (starting with { or [) will be
// parsed and displayed as raw JSON instead of escaped strings.
// Default is true.
func WithParseJSONString(parse bool) Option {
	return func(options *option) {
		options.ParseJSONString = parse
	}
}
