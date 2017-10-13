package cinder_test

import (
	"errors"
	"os"
	"os/exec"
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

func (es *entrySuite) TestEntryWithField() {
	f := es.Entry.WithField("test key", "test value")
	es.Equal("test value", f.Fields["test key"])
}

func (es *entrySuite) TestEntryWithFields() {
	e := es.Entry.WithFields(cinder.Fields{
		"test key 1": "value1",
		"test key 2": 123456,
	})
	es.Len(e.Fields, 2)
	es.Equal("value1", e.Fields["test key 1"])
	es.Equal(123456, e.Fields["test key 2"])
}

func (es *entrySuite) TestEntryWithError() {
	err := errors.New("test error")
	e := es.Entry.WithError(err)
	es.EqualError(err, "test error")
	es.Equal("test error", e.Fields["error"])
}

func (es *entrySuite) TestEntryWithErrorMultiField() {
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

func (es *entrySuite) TestEntryNewTrace() {
	trace := es.Logger.Trace("new trace")
	es.NotNil(trace)
	es.Equal("new trace", trace.Message)
}

func (es *entrySuite) TestEntryTraceStopWithoutError() {
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

func (es *entrySuite) TestEntryTraceStopWithError() {
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

func (es *entrySuite) TestEntryLevel() {
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

func (es *entrySuite) TestEntryDebug() {
	es.Entry.Debug("debug message")

	e := es.Handler.Entries[0]

	es.Exactly("debug message", e.Message)
	es.Exactly(cinder.Level(cinder.DebugLevel), e.Level)
}

func (es *entrySuite) TestEntryFatal() {
	if os.Getenv("ENTRYCRASH") == "1" {
		es.Entry.Fatal("this should crash")
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestEntryFatal")
	cmd.Env = append(os.Environ(), "ENTRYCRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		es.Exactly(cinder.Level(cinder.FatalLevel), es.Entry.Level)
		es.Exactly("this should crash", es.Entry.Message)
	}
}

func (es *entrySuite) TestEntryError() {
	es.Logger.Level = cinder.ErrorLevel
	es.Entry.Error("error message")

	e := es.Handler.Entries[0]

	es.Exactly("error message", e.Message)
	es.Exactly(cinder.Level(cinder.ErrorLevel), e.Level)
}

func (es *entrySuite) TestEntryWarn() {
	es.Logger.Level = cinder.WarnLevel
	es.Entry.Warn("warn message")

	e := es.Handler.Entries[0]

	es.Exactly("warn message", e.Message)
	es.Exactly(cinder.Level(cinder.WarnLevel), e.Level)
}

func (es *entrySuite) TestEntryInfo() {
	es.Logger.Level = cinder.InfoLevel
	es.Entry.Info("info message")

	e := es.Handler.Entries[0]

	es.Exactly("info message", e.Message)
	es.Exactly(cinder.Level(cinder.InfoLevel), e.Level)
}

func (es *entrySuite) TestEntrySilence() {
	es.Logger.Level = cinder.InfoLevel
	es.Entry.Warn("warn message")

	es.Empty(es.Handler.Entries)
}

func TestEntrySuite(t *testing.T) {
	suite.Run(t, new(entrySuite))
}
