package file_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/quinnjr/cinder"
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
	f.TimestampFormat = ""
	f.Format = ""
	fs.Equal("", f.Format)
	fs.Equal("", f.TimestampFormat)
	fs.Implements((*cinder.Handler)(nil), f)
}

func expectedFileContents(e *cinder.Entry) string {
	fc := fmt.Sprintf("[%s] [%s] %s", e.Timestamp.Format("02 Jan 2006 03:04:05 MST"), e.Level, e.Message)
	names := e.Fields.Keys()
	for _, name := range names {
		if name == "source" {
			continue
		}
		fc = fc + fmt.Sprintf("%*s%s=%v", 3, "", name, e.Fields[name])
	}
	fc = fc + fmt.Sprintln()
	return fc
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

	expected := expectedFileContents(entry)

	actual := buf.String()
	fs.NotEmpty(actual)

	fs.Equal(expected, actual)
}
