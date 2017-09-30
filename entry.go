package cinder

import "time"

// Entry ...
type Entry struct {
	Logger    *Logger   `json:"-"`
	Fields    Fields    `json:"fields,omitempty"`
	Level     Level     `json:"level"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

// NewEntry returns a new log entry.
func NewEntry(l *Logger) *Entry {
	return &Entry{Logger: l, Fields: make(map[string]interface{})}
}

// WithFields ...
func (e *Entry) WithFields(fields Fields) *Entry {
	f := e.Fields
	for k, v := range fields {
		f[k] = v
	}
	e.Fields = f
	return e
}

// WithField ...
func (e *Entry) WithField(key string, value interface{}) *Entry {
	field := map[string]interface{}{
		key: value,
	}
	return e.WithFields(field)
}

// Debug ...
func (e *Entry) Debug(msg string) {

}
