package text

import (
	"bytes"
	"testing"

	"github.com/quinnjr/cinder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TextSuite struct {
	suite.Suite
}

func TestTextSuite(t *testing.T) {
	suite.Run(t, new(TextSuite))
}

func (s *TextSuite) TestDefault() {
	h := Default()
	assert.NotNil(s.T(), h, "the handler should be initialized")
	// assert.Implements(t, cinder.Handler, h, "handler should implement the Handler interface")
}

func (s *TextSuite) TestNew() {
	var buf bytes.Buffer
	logger := cinder.New(cinder.DebugLevel, New(&buf))
	assert.NotNil(s.T(), logger, "logger should be initialized")

	// assert.Equal(t, cinder.Logger, logger, "logger should be an instance of cinder.Logger")

}

func (s *TextSuite) TestHandleLog() {
	var buf bytes.Buffer

	logger := cinder.New(cinder.DebugLevel, New(&buf))
	logger.WithField("user", "quinnjr").WithFields(cinder.Fields{
		"test": "1234",
	}).Debug("hello world")

	expected := "[debug][0] hello world              test=1234user=quinnjr\n"

	assert.Equal(s.T(), expected, buf.String())

}
