package logi

type Adapter interface {
	Error(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
}

type Noop struct{}

func (Noop) Error(_ string, _ ...interface{}) {}
func (Noop) Info(_ string, _ ...interface{})  {}
func (Noop) Debug(_ string, _ ...interface{}) {}
func (Noop) Warn(_ string, _ ...interface{})  {}
