package cinder

import stdlog "log"

// Logger ...
type Logger struct {
	Level   Level
	Entry   Entry
	Handler Handler
}

// Prefix returns an Entry with a Prefix field set.
func (l *Logger) Prefix(pre string) *Entry {
	return NewEntry(l).WithPrefix(pre)
}

// Trace ...
func (l *Logger) Trace(msg string) *Entry {
	return NewEntry(l).Trace(msg)
}

// WithField returns a new Entry object with a key-value pair set.
func (l *Logger) WithField(key string, value interface{}) *Entry {
	return NewEntry(l).WithField(key, value)
}

// WithFields ...
func (l *Logger) WithFields(fields Fields) *Entry {
	return NewEntry(l).WithFields(fields)
}

// WithError ...
func (l *Logger) WithError(err error) *Entry {
	if err == nil {
		l.Level = SilentLevel
	}
	return NewEntry(l).WithError(err)
}

// Debug ...
func (l *Logger) Debug(msg string) {
	NewEntry(l).Debug(msg)
}

// Fatal ...
func (l *Logger) Fatal(msg string) {
	NewEntry(l).Fatal(msg)
}

// Error ...
func (l *Logger) Error(msg string) {
	NewEntry(l).Error(msg)
}

// Warn ...
func (l *Logger) Warn(msg string) {
	NewEntry(l).Warn(msg)
}

// Info ...
func (l *Logger) Info(msg string) {
	NewEntry(l).Info(msg)
}

// Debugf ...
func (l *Logger) Debugf() {

}

// Fatalf ...
func (l *Logger) Fatalf() {

}

// Errorf ...
func (l *Logger) Errorf() {

}

// Warnf ...
func (l *Logger) Warnf() {

}

// Infof ...
func (l *Logger) Infof() {

}

func (l *Logger) log(level Level, e *Entry) {
	if level < l.Level || level == SilentLevel {
		return
	}
	if err := l.Handler.HandleLog(e); err != nil {
		stdlog.Printf("eror logging: %s", err)
	}
}
