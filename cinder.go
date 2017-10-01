package cinder

// New creates and returns a new cinder logger utilizing the specified handler and setting the desired log level.
func New(level Level, handler Handler) *Logger {
	return &Logger{
		Level:   level,
		Handler: handler,
	}
}
