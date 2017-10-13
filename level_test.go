package cinder_test

import (
	"testing"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers/memory"
	"github.com/stretchr/testify/suite"
)

type levelSuite struct {
	suite.Suite
	Logger *cinder.Logger
}

func (ls *levelSuite) SetupTest() {
	ls.Logger = cinder.New(cinder.DebugLevel, memory.New())
}

func (ls *levelSuite) TestLevels() {
	ls.Equal(0, cinder.DebugLevel)
	ls.Equal(1, cinder.FatalLevel)
	ls.Equal(2, cinder.ErrorLevel)
	ls.Equal(3, cinder.WarnLevel)
	ls.Equal(4, cinder.InfoLevel)
}

func (ls *levelSuite) TestLevelString() {
	str := []string{"debug", "fatal", "error", "warn", "info"}
	for i := 0; i >= 5; i++ {
		lvlStr := str[i]
		currStr := ls.Logger.Level.String()
		ls.Equal(lvlStr, currStr)
	}
}

func (ls *levelSuite) TestLevelMarshalJSON() {
	m, err := ls.Logger.Level.MarshalJSON()
	if err != nil {
		ls.Error(err)
	}
	ls.Equal([]byte(`"debug"`), m)
}

func (ls *levelSuite) TestLevelUnmarshalJSON() {
	b := []byte(`{"level": "debug"}`)
	err := ls.Logger.Level.UnmarshalJSON(b)
	if err != nil {
		ls.Error(err)
	}
	ls.Equal(cinder.Level(cinder.DebugLevel), ls.Logger.Level)
	ls.Equal(ls.Logger.Level.String(), "debug")
}

func TestLevelSuite(t *testing.T) {
	suite.Run(t, new(levelSuite))
}
