package testify_test

import (
	"testing"

	"github.com/dwhittle/test/testify"
	"github.com/stretchr/testify/suite"
)

type FooTestSuite struct {
	suite.Suite
	start int
}

func (suite *FooTestSuite) SetupTest() {
	suite.start = 5000000
}

func (suite *FooTestSuite) TestExample() {
	n := 1
	if b, ok := suite.T().(*testing.B); ok {
		n = b.N
	}

	for i := 0; i < n; i++ {
		testify.Something(suite.start)
	}
}

func TestFooTestSuite(t *testing.T) {
	suite.Run(t, new(FooTestSuite))
}

func BenchmarkFooTestSuite(b *testing.B) {
	suite.Run(b, new(FooTestSuite))
}
