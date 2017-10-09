package cinder_test

import (
	"errors"
	"testing"
	"time"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers/memory"
	"github.com/stretchr/testify/suite"
)

type entrySuite struct {
	suite.Suite
	Logger  *cinder.Logger
	Entry   *cinder.Entry
	Handler *memory.Handler
}

func (es *entrySuite) SetupTest() {
	es.Handler = memory.New()
	es.Logger = cinder.New(cinder.DebugLevel, es.Handler)
	es.Entry = es.Logger.NewEntry()
}

func (es *entrySuite) TestWithField() {
	f := es.Entry.WithField("test key", "test value")
	es.Equal("test value", f.Fields["test key"])
}

func (es *entrySuite) TestWithFields() {
	e := es.Entry.WithFields(cinder.Fields{
		"test key 1": "value1",
		"test key 2": 123456,
	})
	es.Len(e.Fields, 2)
	es.Equal("value1", e.Fields["test key 1"])
	es.Equal(123456, e.Fields["test key 2"])
}

func (es *entrySuite) TestWithError() {
	err := errors.New("test error")
	e := es.Entry.WithError(err)
	es.EqualError(err, "test error")
	es.Equal("test error", e.Fields["error"])
}

func (es *entrySuite) TestWithErrorMultiField() {
	err := errors.New("test error")
	es.Error(err)
	e := es.Entry.WithError(err).WithFields(cinder.Fields{
		"field1": "bacon ipsum",
		"field2": false,
	})

	es.Len(e.Fields, 3)
	es.Equal("test error", e.Fields["error"])
	es.Equal("bacon ipsum", e.Fields["field1"])
	es.Equal(false, e.Fields["field2"])
}

func (es *entrySuite) TestNewTrace() {
	trace := es.Logger.Trace("new trace")
	es.NotNil(trace)
	es.Equal("new trace", trace.Message)
}

func (es *entrySuite) TestTraceStopWithoutError() {
	func() (err error) {
		defer es.Logger.WithField("file", "wut.png").Trace("test trace").Stop(&err)
		return nil
	}()

	es.Len(es.Handler.Entries, 2)

	{
		e := es.Handler.Entries[0]
		es.Equal("test trace", e.Message)
		es.Equal("wut.png", e.Fields["file"])
	}
	{
		e := es.Handler.Entries[1]
		es.Equal("test trace", e.Message)
		es.Equal("wut.png", e.Fields["file"])
		es.IsType(time.Duration(0), e.Fields["duration"])
	}
}

func (es *entrySuite) TestTraceStopWithError() {
	func() (err error) {
		err = errors.New("test error")
		defer es.Logger.WithField("file", "wut.png").Trace("test trace").Stop(&err)
		return
	}()

	es.Len(es.Handler.Entries, 2)

	{
		e := es.Handler.Entries[0]
		es.Equal("test trace", e.Message)
		es.Equal("wut.png", e.Fields["file"])
	}
	{
		e := es.Handler.Entries[1]
		es.Equal("test trace", e.Message)
		es.Equal("test error", e.Fields["error"])
		es.Equal("wut.png", e.Fields["file"])
		es.IsType(time.Duration(0), e.Fields["duration"])
	}
}

func (es *entrySuite) TestLevel() {
	es.Equal(cinder.Level(cinder.DebugLevel), es.Entry.Level)
	es.Entry.Level = cinder.FatalLevel
	es.Equal(cinder.Level(cinder.FatalLevel), es.Entry.Level)
	es.Entry.Level = cinder.ErrorLevel
	es.Equal(cinder.Level(cinder.ErrorLevel), es.Entry.Level)
	es.Entry.Level = cinder.WarnLevel
	es.Equal(cinder.Level(cinder.WarnLevel), es.Entry.Level)
	es.Entry.Level = cinder.InfoLevel
	es.Equal(cinder.Level(cinder.InfoLevel), es.Entry.Level)
}

func (es *entrySuite) TestDebug() {
	es.Entry.Debug("debug message")

	e := es.Handler.Entries[0]

	es.Exactly("debug message", e.Message)
	es.Exactly(cinder.Level(cinder.DebugLevel), e.Level)
}

func (es *entrySuite) TestFatal() {
	es.Logger.Level = cinder.FatalLevel
	es.Entry.Fatal("fatal message")

	e := es.Handler.Entries[0]

	es.Exactly("fatal message", e.Message)
	es.Exactly(cinder.Level(cinder.FatalLevel), e.Level)
}

func (es *entrySuite) TestError() {
	es.Logger.Level = cinder.ErrorLevel
	es.Entry.Error("error message")

	e := es.Handler.Entries[0]

	es.Exactly("error message", e.Message)
	es.Exactly(cinder.Level(cinder.ErrorLevel), e.Level)
}

func (es *entrySuite) TestWarn() {
	es.Logger.Level = cinder.WarnLevel
	es.Entry.Warn("warn message")

	e := es.Handler.Entries[0]

	es.Exactly("warn message", e.Message)
	es.Exactly(cinder.Level(cinder.WarnLevel), e.Level)
}

func (es *entrySuite) TestInfo() {
	es.Logger.Level = cinder.InfoLevel
	es.Entry.Info("info message")

	e := es.Handler.Entries[0]

	es.Exactly("info message", e.Message)
	es.Exactly(cinder.Level(cinder.InfoLevel), e.Level)
}

func (es *entrySuite) TestSilence() {
	es.Logger.Level = cinder.InfoLevel
	es.Entry.Warn("warn message")

	es.Empty(es.Handler.Entries)
}

func TestEntrySuite(t *testing.T) {
	suite.Run(t, new(entrySuite))
}
