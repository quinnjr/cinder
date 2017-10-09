package cinder

import (
	"bytes"
	"errors"
	"strings"
)

// Log levels.
const (
	silentLevel = iota - 1
	DebugLevel
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
)

// Level of desired log output.
type Level int

// names is an array of strings corresponding to the
// const log levels.
var names = []string{
	DebugLevel: "debug",
	FatalLevel: "fatal",
	ErrorLevel: "error",
	WarnLevel:  "warn",
	InfoLevel:  "info",
}

// levels is a map of the string representation of a level to
// the corresponding log level constant.
var levels = map[string]Level{
	"debug": DebugLevel,
	"fatal": FatalLevel,
	"error": ErrorLevel,
	"warn":  WarnLevel,
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
func (l *Level) UnmarshalJSON(b []byte) (err error) {
	s := string(bytes.Trim(b, `"`))
	v, ok := levels[strings.ToLower(s)]
	if !ok {
		err = errors.New("invalid log level")
	}
	*l = v
	return
}
