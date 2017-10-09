package cinder_test

//
// import (
// 	"testing"
//
// 	"github.com/stretchr/testify/suite"
// )
//
// type fieldsSuite struct {
// 	suite.Suite
// 	Fields Fields
// }
//
// func (fs *fieldsSuite) SetupSuite() {
// 	fs.Fields = Fields{
// 		"test1": 1234567,
// 		"test2": "bacon bacon bacon",
// 	}
// }
//
// type fielder interface {
// 	Keys() []string
// }
//
// func (fs *fieldsSuite) TestKeys() {
// 	var k []string
// 	fs.Implements((*fielder)(nil), fs.Fields)
// 	fs.NotPanics(func() {
// 		k = fs.Fields.Keys()
// 	})
// 	fs.NotEmpty(k)
// 	fs.Exactly(Fields{
// 		"test1": 1234567,
// 		"test2": "bacon bacon bacon",
// 		"test3": false,
// 	}, fs.Fields)
// 	fs.Exactly(3, len(k))
// 	fs.Contains(k, "test1")
// 	fs.Contains(k, "test2")
// 	fs.Contains(k, "test3")
// 	fs.Exactly(k[1], "test1")
// 	fs.Exactly(k[2], "test2")
// 	fs.Exactly(k[3], "test3")
// }
//
// func (fs *fieldsSuite) TestValues() {
// 	var values []interface{}
// 	for _, v := range fs.Fields {
// 		values = append(values, v)
// 	}
// 	fs.Equal(3, values)
// 	fs.Equal(123456, values[1])
// 	fs.Equal("bacon bacon bacon", values[2])
// 	fs.Equal(false, values[3])
// }
//
// func RunFieldSuite(t *testing.T) {
// 	suite.Run(t, new(fieldsSuite))
// }
