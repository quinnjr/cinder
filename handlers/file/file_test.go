package file_test

import (
	"bytes"
	"testing"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers"
	"github.com/quinnjr/cinder/handlers/file"
	"github.com/stretchr/testify/suite"
)

type fileSuite struct {
	suite.Suite
}

func TestFileSuite(t *testing.T) {
	suite.Run(t, new(fileSuite))
}

func (fs *fileSuite) TestNew() {
	var buf bytes.Buffer
	f := file.New(&buf)
	fs.NotNil(f)
	fs.Implements((*cinder.Handler)(nil), f)
}

func (fs *fileSuite) TestNewWithFormat() {
	var buf bytes.Buffer
	f := file.New(&buf)
	fs.NotNil(f)
	fs.Exactly(handlers.DefaultFormat, f.GetFormat())
	fs.Exactly(handlers.DefaultTimestamp, f.GetTimestamp())
	f.SetTimestamp("")
	f.SetFormat("")
	fs.Equal("", f.GetFormat())
	fs.Equal("", f.GetTimestamp())
	fs.Implements((*cinder.Handler)(nil), f)
}

func (fs *fileSuite) TestHandleLog() {
	var buf bytes.Buffer
	h := file.New(&buf)
	logger := cinder.New(cinder.DebugLevel, h)
	entry := logger.WithFields(cinder.Fields{
		"test1":  "key1",
		"source": "galapagos",
	})
	entry.Debug("test log entry")

	expected := "[DEBUG] [" + entry.Timestamp.Format(handlers.DefaultTimestamp) + "] test log entry   test1=key1\n"

	actual := buf.String()
	fs.NotEmpty(actual)

	fs.Equal(expected, actual)
}
