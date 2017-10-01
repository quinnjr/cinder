package cinder_test

import (
	"errors"
	"testing"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoggerSuite struct {
	suite.Suite
	Logger *cinder.Logger
}

func (s *LoggerSuite) SetupTest() {
	s.Logger = cinder.New(cinder.DebugLevel, memory.New())
}

func (s *LoggerSuite) TestNew() {
	t := s.T()
	assert.NotNil(t, s.Logger, "Logger should be instantiated")
	assert.Equal(t, &cinder.Logger{Level: cinder.DebugLevel, Handler: memory.New()}, s.Logger, "s.Logger should be a proper instance of Logger")
}

func (s *LoggerSuite) TestPrefix() {
	t := s.T()
	e := s.Logger.Prefix("test")
	assert.NotNil(t, e, "Entry should be properly returned from the Prefix function")
	assert.Equal(t, cinder.Level(0), e.Level, "Entry should be set at level `cinder.DebugLevel`")
	assert.Equal(t, "test", e.Prefix, "Prefix should be `test`")
}

func (s *LoggerSuite) TestTrace() {
	t := s.T()
	trace := s.Logger.Trace("test")
	assert.NotNil(t, trace, "The trace entry should be properly initialized")
	assert.Equal(t, cinder.Level(0), trace.Level, "Trace entry shouuld be set at level `cinder.DebugLevel`")
	assert.Equal(t, "test", trace.Message, "Trace's message should be `test`")
}

func (s *LoggerSuite) TestWithField() {

}

func (s *LoggerSuite) TestWithFields() {

}

func (s *LoggerSuite) TestWithError() {
	t := s.T()
	expected := "Spicy jalapeno bacon ipsum dolor amet boudin shank porchetta tri-tip"
	e := s.Logger.WithError(errors.New(expected))
	assert.NotNil(t, e, "the error log should be initialized")
	assert.Equal(t, expected, e.Fields["error"], "the error field should be equal to the expected text")
}

func (s *LoggerSuite) TestWithError_Nil() {
	t := s.T()
	e := s.Logger.WithError(nil)
	assert.NotNil(t, e, "The error log should be properly initialized")
	assert.Equal(t, cinder.DebugLevel, int(e.Level), "The error log should be set to `SilentLevel` to not actually send out any logging information since `(err error)` is nil")
}

// // Debug ...
// func (l *Logger) Debug(msg string) {
// 	NewEntry(l).Debug(msg)
// }
//
// // Fatal ...
// func (l *Logger) Fatal(msg string) {
// 	NewEntry(l).Fatal(msg)
// }
//
// // Error ...
// func (l *Logger) Error(msg string) {
// 	NewEntry(l).Error(msg)
// }
//
// // Warn ...
// func (l *Logger) Warn(msg string) {
// 	NewEntry(l).Warn(msg)
// }
//
// // Info ...
// func (l *Logger) Info(msg string) {
// 	NewEntry(l).Info(msg)
// }
//
// // Debugf ...
// func (l *Logger) Debugf() {
//
// }
//
// // Fatalf ...
// func (l *Logger) Fatalf() {
//
// }
//
// // Errorf ...
// func (l *Logger) Errorf() {
//
// }
//
// // Warnf ...
// func (l *Logger) Warnf() {
//
// }
//
// // Infof ...
// func (l *Logger) Infof() {
//
// }
//
// func (l *Logger) log(level Level, e *Entry) {
// 	if level < l.Level || level == SilentLevel {
// 		return
// 	}
// 	if err := l.Handler.HandleLog(e); err != nil {
// 		stdlog.Printf("eror logging: %s", err)
// 	}
// }

func TestLoggerSuite(t *testing.T) {
	suite.Run(t, new(LoggerSuite))
}
