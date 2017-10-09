package text_test

import (
	"bytes"
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

	expected := "[DEBUG] [0000] hello world         test=1234user=quinnjr\n"

	ts.Equal(expected, buf.String())

}

func TestTextSuite(t *testing.T) {
	suite.Run(t, new(TextSuite))
}
