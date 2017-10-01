package cinder_test

import (
	"testing"

	"github.com/quinnjr/cinder"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	logger := cinder.New(cinder.DebugLevel, nil)
	assert.NotNil(t, logger, "logger should be initialized")
	assert.Equal(t, cinder.DebugLevel, int(logger.Level), "level should be properly set")
}
