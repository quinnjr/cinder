package cinder

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Entry ...
type Entry struct {
	Logger    *Logger   `json:"-"`
	Fields    Fields    `json:"fields,omitempty"`
	Level     Level     `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	start     time.Time
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// NewEntry returns a new log entry.
func newEntry(l *Logger) *Entry {
	return &Entry{Logger: l, Fields: make(Fields), Level: l.Level}
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
	if err == nil {
		return e
	}

	f := newEntry(e.Logger).WithField("error", err.Error())

	// stack tracing
	if s, ok := err.(stackTracer); ok {
		fr := s.StackTrace()[0]
		name := fmt.Sprintf("%n", fr)
		file := fmt.Sprintf("%+s", fr)
		line := fmt.Sprintf("%d", fr)

		p := strings.Split(file, "\n\t")
		if len(p) > 1 {
			file = p[1]
		}
		f = f.WithField("source", fmt.Sprintf("%s: %s:%s", name, file, line))
	}
	// Merge already set Fields with the error field.
	f = f.WithFields(e.Fields)
	return f
}

// Trace ...
func (e *Entry) Trace(msg string) *Entry {
	e.Info(msg)
	tr := e.WithFields(e.Fields)
	tr.Message = msg
	tr.start = time.Now()
	return tr
}

// Stop ...
func (e *Entry) Stop(err *error) {
	if err == nil || *err == nil {
		e.WithField("duration", time.Since(e.start)).Info(e.Message)
	} else {
		e.WithError(*err).WithField("duration", time.Since(e.start)).Error(e.Message)
	}
}

// Debug ...
func (e *Entry) Debug(msg string) error {
	return e.log(DebugLevel, msg)
}

// Fatal ...
func (e *Entry) Fatal(msg string) (err error) {
	err = e.log(FatalLevel, msg)
	if err == nil {
		os.Exit(1)
	}
	return
}

// Error ...
func (e *Entry) Error(msg string) error {
	return e.log(ErrorLevel, msg)
}

// Warn ...
func (e *Entry) Warn(msg string) error {
	return e.log(WarnLevel, msg)
}

// Info ...
func (e *Entry) Info(msg string) error {
	return e.log(InfoLevel, msg)
}

// Debugf ...
func (e *Entry) Debugf(format string, a ...interface{}) error {
	return e.Debug(fmt.Sprintf(format, a))
}

// Fatalf ...
func (e *Entry) Fatalf(format string, a ...interface{}) error {
	return e.Fatal(fmt.Sprintf(format, a))
}

// Errorf ...
func (e *Entry) Errorf(format string, a ...interface{}) error {
	return e.Error(fmt.Sprintf(format, a))
}

// Warnf ...
func (e *Entry) Warnf(format string, a ...interface{}) error {
	return e.Warn(fmt.Sprintf(format, a))
}

// Infof ...
func (e *Entry) Infof(format string, a ...interface{}) error {
	return e.Info(fmt.Sprintf(format, a))
}

// finalize returns the Entry with all of the correct fields populated.
func (e *Entry) finalize(level Level, msg string) *Entry {
	e.Level = level
	e.Message = msg
	e.Timestamp = time.Now()
	return e
}

// log handlers the actual logging event by checking if the
// argument level is above the threshold of the Logger's
// Level field. If the level is good, log calls the Handler's
// HandlerLog fucntion with the finalized level and message.
func (e *Entry) log(level Level, msg string) (err error) {
	l := e.Logger
	if level < l.Level || level == silentLevel {
		return
	}
	err = l.Handler.HandleLog(e.finalize(level, msg))
	return
}
