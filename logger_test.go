package cinder_test

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers/memory"
	"github.com/stretchr/testify/suite"
)

type loggerSuite struct {
	suite.Suite
	Logger  *cinder.Logger
	Handler *memory.Handler
}

func (ls *loggerSuite) SetupTest() {
	ls.Handler = memory.New()
	ls.Logger = cinder.New(cinder.DebugLevel, ls.Handler)
}

func (ls *loggerSuite) TestLoggerNew() {
	ls.NotNil(ls.Logger)
	ls.Equal(&cinder.Logger{Level: cinder.DebugLevel, Handler: memory.New()}, ls.Logger)
}

func (ls *loggerSuite) TestLoggerWithField() {
	e := ls.Logger.WithField("hello", "world")
	ls.Len(e.Fields, 1)
	ls.Equal("world", e.Fields["hello"])
}

func (ls *loggerSuite) TestLoggerWithFields() {
	f := cinder.Fields{
		"hello": "world",
		"1234":  1234,
	}
	e := ls.Logger.WithFields(f)
	ls.Len(e.Fields, 2)
	ls.Equal(f, e.Fields)
}

func (ls *loggerSuite) TestLoggerWithError() {
	expected := "Spicy jalapeno bacon ipsum dolor amet boudin shank porchetta tri-tip"
	e := ls.Logger.WithError(errors.New(expected))
	ls.NotNil(e)
	ls.Equal(expected, e.Fields["error"])
}

func (ls *loggerSuite) TestLoggerWithError_Nil() {
	e := ls.Logger.WithError(nil)
	ls.NotNil(e)
	ls.Equal(cinder.Level(-1), e.Level)
}

func (ls *loggerSuite) TestLoggerDebug() {
	f := cinder.Fields{"test1": "value1"}
	ls.Logger.WithFields(f).Debug("debug message")
	e := ls.Handler.Entries[0]

	ls.Exactly("debug message", e.Message)
	ls.Exactly(f, e.Fields)
}

func (ls *loggerSuite) TestLoggerFatal() {
	f := cinder.Fields{"error": "test error"}
	if os.Getenv("LOGGERCRASH") == "1" {
		ls.Logger.WithError(errors.New("test error")).Fatal("this should crash")
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLoggerFatal")
	cmd.Env = append(os.Environ(), "LOGGERCRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		ls.Exactly(cinder.Level(cinder.FatalLevel), ls.Logger.Level)
		ls.Exactly("this should crash", ls.Logger.Entry.Message)
		ls.Exactly(f, ls.Logger.Entry.Fields)
	}
}

func (ls *loggerSuite) TestLoggerError() {
	ls.Logger = cinder.New(cinder.ErrorLevel, ls.Handler)
	err := errors.New("test error")
	f := cinder.Fields{"error": "test error"}
	ls.Logger.WithError(err).Error("error message")
	e := ls.Handler.Entries[0]

	ls.Exactly(cinder.Level(cinder.ErrorLevel), e.Level)
	ls.Exactly("error message", e.Message)
	ls.Exactly(f, e.Fields)
}

func (ls *loggerSuite) TestLoggerWarn() {
	ls.Logger = cinder.New(cinder.WarnLevel, ls.Handler)
	f := cinder.Fields{"warn field": "warn value"}
	ls.Logger.WithFields(f).Warn("warn message")
	e := ls.Handler.Entries[0]

	ls.Exactly(cinder.Level(cinder.WarnLevel), e.Level)
	ls.Exactly("warn message", e.Message)
	ls.Exactly(f, e.Fields)
}

func (ls *loggerSuite) TestLoggerInfo() {
	ls.Logger = cinder.New(cinder.InfoLevel, ls.Handler)
	f := cinder.Fields{"info field": "info value"}
	ls.Logger.WithFields(f).Info("info message")
	e := ls.Handler.Entries[0]

	ls.Exactly(cinder.Level(cinder.InfoLevel), e.Level)
	ls.Exactly("info message", e.Message)
	ls.Exactly(f, e.Fields)
}

func (ls *loggerSuite) TestLoggerDebugf() {
	f := cinder.Fields{"debug field": "debug value"}
	ls.Logger.WithFields(f).Debugf("%s", "debug message")
	e := ls.Handler.Entries[0]

	ls.Exactly(cinder.Level(cinder.DebugLevel), e.Level)
	ls.Exactly("[debug message]", e.Message)
	ls.Exactly(f, e.Fields)
}

func (ls *loggerSuite) TestLoggerFatalf() {
	if os.Getenv("LOGGERCRASH") == "2" {
		ls.Logger.Fatalf("%s", "this should crash")
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLoggerFatalf")
	cmd.Env = append(os.Environ(), "LOGGERCRASH=2")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		ls.Exactly(cinder.Level(cinder.FatalLevel), ls.Logger.Level)
		ls.Exactly("this should crash", ls.Logger.Entry.Message)
	}
}

func (ls *loggerSuite) TestLoggerErrorf() {
	err := errors.New("error error")
	f := cinder.Fields{"error": "error error"}
	ls.Logger.Level = cinder.ErrorLevel
	ls.Logger.WithError(err).Errorf("%s", "error message")
	e := ls.Handler.Entries[0]

	ls.Exactly(cinder.Level(cinder.ErrorLevel), e.Level)
	ls.Exactly("[error message]", e.Message)
	ls.Exactly(f, e.Fields)
}

func (ls *loggerSuite) TestLoggerWarnf() {
	f := cinder.Fields{"warning": "warning thing"}
	ls.Logger.Level = cinder.WarnLevel
	ls.Logger.WithFields(f).Warnf("%s", "warn message")
	e := ls.Handler.Entries[0]

	ls.Exactly(cinder.Level(cinder.WarnLevel), e.Level)
	ls.Exactly("[warn message]", e.Message)
	ls.Exactly(f, e.Fields)
}

func (ls *loggerSuite) TestLoggerInfof() {
	f := cinder.Fields{"info": "info thing"}
	ls.Logger.Level = cinder.InfoLevel
	ls.Logger.WithFields(f).Infof("%s", "info message")
	e := ls.Handler.Entries[0]

	ls.Exactly(cinder.Level(cinder.InfoLevel), e.Level)
	ls.Exactly("[info message]", e.Message)
	ls.Exactly(f, e.Fields)
}

func TestLoggerSuite(t *testing.T) {
	suite.Run(t, new(loggerSuite))
}
