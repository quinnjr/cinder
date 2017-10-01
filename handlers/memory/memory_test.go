package memory_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MemorySuite struct {
	suite.Suite
}

func TestMemorySuite(t *testing.T) {
	suite.Run(t, new(MemorySuite))
}
