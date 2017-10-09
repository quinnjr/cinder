package memory_test

import (
	"testing"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers/memory"
	"github.com/stretchr/testify/suite"
)

type MemorySuite struct {
	suite.Suite
	Handler *memory.Handler
}

func (ms *MemorySuite) TestNew() {
	ms.NotPanics(func() {
		ms.Handler = memory.New()
	})

	ms.NotEmpty(ms.Handler)
	ms.Implements((*cinder.Handler)(nil), ms.Handler)
}

func (ms *MemorySuite) TestHandleLog() {
	ms.NotPanics(func() {
		ms.Handler = memory.New()
	})

	logger := cinder.New(cinder.DebugLevel, ms.Handler)
	e := logger.WithField("test1", 123)
	e.Info("test message")
	ms.Equal("test message", e.Message)

	ms.NotEmpty(ms.Handler.Entries)
	ms.Len(ms.Handler.Entries, 1)
	ms.Equal(e, ms.Handler.Entries[0])
}

func TestMemorySuite(t *testing.T) {
	suite.Run(t, new(MemorySuite))
}
