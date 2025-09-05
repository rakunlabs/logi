package logadapter

type Adapter interface {
	Error(msg string, keysAndValues ...any)
	Info(msg string, keysAndValues ...any)
	Debug(msg string, keysAndValues ...any)
	Warn(msg string, keysAndValues ...any)
}

type Noop struct{}

func (Noop) Error(_ string, _ ...any) {}
func (Noop) Info(_ string, _ ...any)  {}
func (Noop) Debug(_ string, _ ...any) {}
func (Noop) Warn(_ string, _ ...any)  {}
