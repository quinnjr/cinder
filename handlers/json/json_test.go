package json_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/quinnjr/cinder"
	jHandler "github.com/quinnjr/cinder/handlers/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JSONSuite struct {
	suite.Suite
}

func TestJSONSuite(t *testing.T) {
	suite.Run(t, new(JSONSuite))
}

func (s *JSONSuite) SetupSuite() {

}

func (s *JSONSuite) TestNew() {
	var buf bytes.Buffer
	h := jHandler.New(&buf)
	assert.NotNil(s.T(), h, "Handler should be properly instantiated")
	assert.Implements(s.T(), (*cinder.Handler)(nil), h, "New should return a Handler which implements the Handler interface")

	logger := cinder.New(cinder.DebugLevel, h)
	assert.NotNil(s.T(), logger, "Logger instance should be properly instantiated with the json handler")

	entry := logger.WithField("user", "joe")
	entry.Info("hello world")

	json, err := json.Marshal(entry)
	if err != nil {
		s.Error(err)
	}

	expected := string(json[:]) + "\n"

	assert.Equal(s.T(), expected, buf.String(), "Output from the logger should be contain the expected output")
}

func (s *JSONSuite) TestDefault() {
	h := jHandler.Default()
	assert.NotNil(s.T(), h, "Handler should be properly instantiated")
	assert.Implements(s.T(), (*cinder.Handler)(nil), h, "New should return a Handler which implements the Handler interface")

	logger := cinder.New(cinder.DebugLevel, h)
	assert.NotNil(s.T(), logger, "Logger instance should be properly instantiated with the default json handler")
}
