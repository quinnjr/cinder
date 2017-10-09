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

func (js *JSONSuite) SetupSuite() {

}

func (js *JSONSuite) TestNew() {
	var buf bytes.Buffer
	t := js.T()
	h := jHandler.New(&buf)
	assert.NotNil(t, h, "Handler should be properly instantiated")
	assert.Implements(t, (*cinder.Handler)(nil), h, "New should return a Handler which implements the Handler interface")

	logger := cinder.New(cinder.DebugLevel, h)
	assert.NotNil(t, logger, "Logger instance should be properly instantiated with the json handler")

	entry := logger.WithField("user", "joe")
	entry.Info("hello world")

	json, err := json.Marshal(entry)
	if err != nil {
		js.Error(err)
	}

	expected := string(json[:]) + "\n"

	assert.Equal(t, expected, buf.String(), "Output from the logger should be contain the expected output")
}

func (js *JSONSuite) TestDefault() {
	t := js.T()
	h := jHandler.Default()
	assert.NotNil(t, h, "Handler should be properly instantiated")
	assert.Implements(t, (*cinder.Handler)(nil), h, "New should return a Handler which implements the Handler interface")

	logger := cinder.New(cinder.DebugLevel, h)
	assert.NotNil(t, logger, "Logger instance should be properly instantiated with the default json handler")
}
