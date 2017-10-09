package cinder_test

import (
	"testing"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers/memory"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	l := cinder.New(cinder.DebugLevel, nil)
	assert.NotNil(t, l)
	assert.Equal(t, cinder.Level(cinder.DebugLevel), l.Level)
}

func TestNewMemory(t *testing.T) {
	l := cinder.New(cinder.DebugLevel, memory.New())
	assert.NotNil(t, l)
	assert.Implements(t, (*cinder.Handler)(nil), l.Handler)
	assert.Equal(t, cinder.Level(cinder.DebugLevel), l.Level)
}
