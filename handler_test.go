package cinder_test

import (
	"testing"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers/memory"
	"github.com/stretchr/testify/suite"
)

type handlerSuite struct {
	suite.Suite
	Logger  *cinder.Logger
	Handler *memory.Handler
}

func (hs *handlerSuite) SetupTest() {
	hs.Handler = memory.New()
	hs.Logger = cinder.New(cinder.DebugLevel, hs.Handler)
}

func (hs *handlerSuite) TestHandlerInterface() {
	hs.Implements((*cinder.Handler)(nil), hs.Handler)
}

func (hs *handlerSuite) TestHandlerHandleLog() {
	e := hs.Logger.WithField("test", "field")
	err := hs.Handler.HandleLog(e)
	hs.NoError(err)
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}
