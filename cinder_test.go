package cinder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	logger := New(DebugLevel, nil)
	assert.NotNil(t, logger, "logger should be initialized")
}
