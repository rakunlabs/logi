package logi

import (
	"io"
	"os"
	"strconv"

	"github.com/mattn/go-isatty"
)

type selection uint8

const (
	SelectAuto selection = iota
	SelectTrue
	SelectFalse
)

type fd interface {
	Fd() uintptr
}

func prettySelection(v string) selection {
	if v == "" {
		return SelectAuto
	}

	vBool, _ := strconv.ParseBool(v)
	if vBool {
		return SelectTrue
	}

	return SelectFalse
}

func isPretty(v selection, w io.Writer) bool {
	switch v {
	case SelectAuto:
		v, ok := os.LookupEnv(EnvPretty)
		if ok {
			result, _ := strconv.ParseBool(v)

			return result
		}

		if w, ok := w.(fd); ok {
			return isatty.IsTerminal(w.Fd()) || isatty.IsCygwinTerminal(w.Fd())
		}

		return false
	case SelectFalse:
		return false
	case SelectTrue:
		return true
	}

	return false
}

func checkLevel(level string) string {
	if v, ok := os.LookupEnv(EnvLevel); ok {
		return v
	}

	return level
}

func timeFormat(opt *string, pretty bool) string {
	if opt != nil {
		return *opt
	}

	if pretty {
		return TimePrettyFormat
	}

	return TimeFormat
}
