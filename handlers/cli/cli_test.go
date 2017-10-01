package cli_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CLISuite struct {
	suite.Suite
}

func TestCLISuite(t *testing.T) {
	suite.Run(t, new(CLISuite))
}
