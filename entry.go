package cinder

import (
	"fmt"
	"time"
)

// Entry ...
type Entry struct {
	Logger    *Logger   `json:"-"`
	Fields    Fields    `json:"fields,omitempty"`
	Level     Level     `json:"level"`
	Message   string    `json:"message"`
	Prefix    string    `json:"prefix,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	start     time.Time
}

// NewEntry returns a new log entry.
func NewEntry(l *Logger) *Entry {
	return &Entry{Logger: l, Fields: Fields{}, Timestamp: time.Now().Local(), start: time.Now().Local()}
}

// WithPrefix returns an Entry with a prefix.
func (e *Entry) WithPrefix(pre string) *Entry {
	e.Prefix = pre
	return e
}

// WithFields returns an Entry with the supplied Fields added to the Entry's Fields.
func (e *Entry) WithFields(fields Fields) *Entry {
	f := e.Fields
	for k, v := range fields {
		f[k] = v
	}
	e.Fields = f
	return e
}

// WithField returns an Entry with the supplied Field added to the Entry's Fields.
func (e *Entry) WithField(key string, value interface{}) *Entry {
	return e.WithFields(Fields{key: value})
}

// WithError returns an Entry with err set as a Field
func (e *Entry) WithError(err error) *Entry {
	f := e.Fields
	// Ensure that the "error" field is the first field if a valid error is recieved.
	var m Fields
	if err != nil {
		m = Fields{"error": err.Error()}
	}
	// TODO: Implement stack trace here
	// Merge already set Fields with the error field.
	for k := range f {
		m[k] = f[k]
	}
	e.Fields = m
	return e
}

// Trace ...
func (e *Entry) Trace(msg string) *Entry {
	e.Info(msg)
	return &Entry{
		Logger:    e.Logger,
		Level:     e.Level,
		Fields:    e.Fields,
		Timestamp: e.Timestamp,
		Message:   e.Message,
		Prefix:    e.Prefix,
		start:     time.Now(),
	}
}

// Stop ...
func (e *Entry) Stop(err *error) {
	if err == nil || *err == nil {
		e.WithField("duration", time.Since(e.start)).Info(e.Message)
	} else {
		e.WithField("duration", time.Since(e.start)).WithError(*err).Error(e.Message)
	}
}

// Debug ...
func (e *Entry) Debug(msg string) {
	e.Message = msg
	e.Logger.log(DebugLevel, e)
}

// Fatal ...
func (e *Entry) Fatal(msg string) {
	e.Message = msg
	e.Logger.log(FatalLevel, e)
}

// Error ...
func (e *Entry) Error(msg string) {
	e.Message = msg
	e.Logger.log(ErrorLevel, e)
}

// Warn ...
func (e *Entry) Warn(msg string) {
	e.Message = msg
	e.Logger.log(WarnLevel, e)
}

// Info ...
func (e *Entry) Info(msg string) {
	e.Message = msg
	e.Logger.log(InfoLevel, e)
}

// Debugf ...
func (e *Entry) Debugf(format string, a ...interface{}) {
	e.Debug(fmt.Sprintf(format, a))
}

// Fatalf ...
func (e *Entry) Fatalf(format string, a ...interface{}) {
	e.Fatal(fmt.Sprintf(format, a))
}

// Errorf ...
func (e *Entry) Errorf(format string, a ...interface{}) {
	e.Error(fmt.Sprintf(format, a))
}

// Warnf ...
func (e *Entry) Warnf(format string, a ...interface{}) {
	e.Warn(fmt.Sprintf(format, a))
}

// Infof ...
func (e *Entry) Infof(format string, a ...interface{}) {
	e.Info(fmt.Sprintf(format, a))
}
