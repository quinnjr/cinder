package cinder

// Logger ...
type Logger struct {
	Level   Level
	Entry   Entry
	Handler *Handler
}

// New creates and returns a new logger utilizing the specified handler and setting the desired log level.
func New(level Level, handler *Handler) *Logger {
	return &Logger{
		Level:   level,
		Handler: handler,
	}
}
