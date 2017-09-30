package cinder

import (
	"bytes"
	"errors"
	"strings"
)

// Log levels.
const (
	DebugLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
)

// Level of desired log output.
type Level int

var names = []string{
	DebugLevel: "debug",
	FatalLevel: "fatal",
	ErrorLevel: "error",
	WarnLevel:  "warn",
	InfoLevel:  "info",
}

var levels = map[string]Level{
	"debug": DebugLevel,
	"fatal": FatalLevel,
	"error": ErrorLevel,
	"info":  InfoLevel,
}

// String returns the string representation of the log level.
func (l Level) String() string {
	return names[l]
}

// MarshalJSON implements the json.Marshaler interface.
func (l Level) MarshalJSON() ([]byte, error) {
	return []byte(`"` + l.String() + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (l *Level) UnmarshalJSON(b []byte) error {
	s := string(bytes.Trim(b, `"`))
	v, ok := levels[strings.ToLower(s)]
	if !ok {
		return errors.New("invalid log level")
	}
	*l = v
	return nil
}
