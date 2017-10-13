package text_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers/text"
	"github.com/stretchr/testify/suite"
)

type TextSuite struct {
	suite.Suite
	Handler *text.Handler
}

func (ts *TextSuite) TearDownSuite() {
	os.Remove("./test.log")
}

func (ts *TextSuite) TestDefault() {
	ts.NotPanics(func() {
		ts.Handler = text.Default()
	})
	ts.NotNil(ts.Handler)
	ts.Equal(os.Stderr, ts.Handler.Writer)
}

func (ts *TextSuite) TestNew() {
	ts.NotPanics(func() {
		ts.Handler = text.New(os.Stderr)
	})
	ts.NotNil(ts.Handler)
	ts.Equal(os.Stderr, ts.Handler.Writer)
}

func (ts *TextSuite) TestHandleLog() {
	var buf bytes.Buffer

	ts.NotPanics(func() {
		ts.Handler = text.New(&buf)
	})

	logger := cinder.New(cinder.DebugLevel, ts.Handler)
	e := logger.WithField("user", "quinnjr").WithFields(cinder.Fields{
		"test": "1234",
	})
	e.Debug("hello world")

	expected := "[DEBUG] [" + e.Timestamp.Format(text.DefaultTimestamp) + "] hello world            test=1234   user=quinnjr\n"

	ts.Equal(expected, buf.String())
}

func (ts *TextSuite) TestFileOutput() {
	file, err := os.Create("./test.log")
	ts.NoError(err)
	defer file.Close()

	logger := cinder.New(cinder.DebugLevel, text.New(file))

	e := logger.WithField("test key", "test value")

	err = e.Info("test logging")
	ts.NoError(err)

	b, err := ioutil.ReadFile("./test.log")
	ts.NoError(err)

	expected := "[INFO] [" + e.Timestamp.Format(ts.Handler.GetTimestamp()) + "] test logging            test key=test value\n"

	ts.Exactly(expected, string(b))
}

func (ts *TextSuite) TestSetandGetFormatandTimestamp() {
	ts.Exactly(text.DefaultFormat, ts.Handler.GetFormat())
	ts.Exactly(text.DefaultTimestamp, ts.Handler.GetTimestamp())
	ts.Handler.SetFormat("")
	ts.Handler.SetTimestamp("")
	ts.Exactly("", ts.Handler.GetFormat())
	ts.Exactly("", ts.Handler.GetTimestamp())
}

func TestTextSuite(t *testing.T) {
	suite.Run(t, new(TextSuite))
}
