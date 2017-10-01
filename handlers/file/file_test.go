package file_test

import (
	"os"
	"testing"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FileSuite struct {
	suite.Suite
}

func TestFileSuite(t *testing.T) {
	suite.Run(t, new(FileSuite))
}

func (s *FileSuite) SetupSuite() {
	if _, err := os.Stat("test.log"); !os.IsExist(err) {
		if _, err := os.Create("test.log"); err != nil {
			s.Error(err)
		}
	}
}

func (s *FileSuite) TearDownSuite() {
	if err := os.Remove("test.log"); err != nil {
		s.Error(err)
	}
}

func (s *FileSuite) TestNew() {
	t := s.T()
	f := file.New("./test.log", nil, nil)
	assert.NotNil(t, f, "A handler should be returned from file.New")
	assert.Equal(t, "./test.log", f.Filepath, "Filepath should be set to the proper value")
	assert.Equal(t, file.DefaultLogFormat, f.Format, "Format should be set to the default format")
	assert.Equal(t, file.DefaultTimestampFormat, f.Timestamp, "Timestamp should be set to the default timestamp format")
	assert.Implements(t, (*cinder.Handler)(nil), f, "New should return a Handler which implements the Handler interface")
}

func (s *FileSuite) TestNewWithFormat() {
	t := s.T()
	f := file.New("./test.log", "", "")
	assert.NotNil(t, f, "A handler should be returned from file.New")
	assert.Equal(t, "./test.log", f.Filepath, "Filepath should be set to the proper value")
	assert.Equal(t, "", f.Format, "Format should be set to the default format")
	assert.Equal(t, "", f.Timestamp, "Timestamp should be set to the default timestamp format")
	assert.Implements(t, (*cinder.Handler)(nil), f, "New should return a Handler which implements the Handler interface")
}

func (s *FileSuite) TestNewWithPanic() {
	t := s.T()
	assert.Panics(t, func() {
		file.New("./test.log", []byte("fail"), nil)
	})
	assert.Panics(t, func() {
		file.New("./test.log", "", []byte("fail"))
	})
}
