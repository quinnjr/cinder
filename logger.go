package cinder

// Cinder ...
type Cinder interface {
	WithFields(fields Fields) *Entry
	WithField(key string, value interface{}) *Entry
	WithError(err error) *Entry
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
}

// Logger ...
type Logger struct {
	Level   Level
	Entry   Entry
	Handler Handler
}

// NewEntry returns a new Entry with the logger set to the logger instance.
func (l *Logger) NewEntry() *Entry {
	return newEntry(l)
}

// Trace ...
func (l *Logger) Trace(msg string) *Entry {
	return newEntry(l).Trace(msg)
}

// WithField returns a new Entry object with a key-value pair set.
func (l *Logger) WithField(key string, value interface{}) *Entry {
	return newEntry(l).WithField(key, value)
}

// WithFields ...
func (l *Logger) WithFields(fields Fields) *Entry {
	return newEntry(l).WithFields(fields)
}

// WithError ...
func (l *Logger) WithError(err error) *Entry {
	if err == nil {
		l.Level = silentLevel
	}
	return newEntry(l).WithError(err)
}

// Debug ...
func (l *Logger) Debug(msg string) {
	newEntry(l).Debug(msg)
}

// Fatal ...
func (l *Logger) Fatal(msg string) {
	newEntry(l).Fatal(msg)
}

// Error ...
func (l *Logger) Error(msg string) {
	newEntry(l).Error(msg)
}

// Warn ...
func (l *Logger) Warn(msg string) {
	newEntry(l).Warn(msg)
}

// Info ...
func (l *Logger) Info(msg string) {
	newEntry(l).Info(msg)
}

// Debugf ...
func (l *Logger) Debugf(format string, v ...interface{}) {
	newEntry(l).Debugf(format, v)
}

// Fatalf ...
func (l *Logger) Fatalf(format string, v ...interface{}) {
	newEntry(l).Fatalf(format, v)
}

// Errorf ...
func (l *Logger) Errorf(format string, v ...interface{}) {
	newEntry(l).Errorf(format, v)
}

// Warnf ...
func (l *Logger) Warnf(format string, v ...interface{}) {
	newEntry(l).Errorf(format, v)
}

// Infof ...
func (l *Logger) Infof(format string, v ...interface{}) {
	newEntry(l).Infof(format, v)
}
