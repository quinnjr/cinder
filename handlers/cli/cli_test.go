package cli_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers/cli"
	"github.com/stretchr/testify/suite"
)

type CLISuite struct {
	suite.Suite
	Handler *cli.Handler
}

func (cl *CLISuite) TestNew() {
	cl.NotPanics(func() {
		cl.Handler = cli.New(os.Stdout)
	})
	cl.Implements((*cinder.Handler)(nil), cl.Handler)
	cl.Equal(os.Stdout, cl.Handler.Writer)
	cl.Equal(3, cl.Handler.Padding)
}

func (cl *CLISuite) TestDefault() {
	cl.NotPanics(func() {
		cl.Handler = cli.Default()
	})
	cl.Implements((*cinder.Handler)(nil), cl.Handler)
	cl.Equal(os.Stderr, cl.Handler.Writer)
	cl.Equal(3, cl.Handler.Padding)
}

func (cl *CLISuite) TestHandleLog() {
	var buf bytes.Buffer
	cl.NotPanics(func() {
		cl.Handler = cli.New(&buf)
	})

	logger := cinder.New(cinder.DebugLevel, cl.Handler)
	e := logger.WithFields(cinder.Fields{
		"test1":  "value1",
		"source": false,
	})
	e.Info("test message")

	expected := e.Timestamp.Format(cli.DefaultTimestamp) + " \x1b[34mINFO\x1b[0m test message        \x1b[34mtest1\x1b[0m=value1    \n"

	cl.Equal(expected, buf.String())
}

func TestCLISuite(t *testing.T) {
	suite.Run(t, new(CLISuite))
}
