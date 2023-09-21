package intTest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go_ether/unit_test/testcase"
	"testing"
)

type SplitTest struct {
	suite.Suite
}

func TestSplit(t *testing.T) {
	suite.Run(t, new(SplitTest))
}

func (s *SplitTest) SetupSuite() {
	fmt.Println("this is SetupSuite")
}

func (s *SplitTest) BeforeTest(suiteName, testName string) {
	fmt.Printf("this is BeforeTest,%s==>%s", suiteName, testName)
}

func (s *SplitTest) Test_00_demo() {
	type args struct {
		s   string
		sep string
	}

	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{},
		{"test1", args{"a:b:c", ":"}, []string{"a", "b", "c"}},
		{"test2", args{"a:b:c", ","}, []string{"a:b:c"}},
		{"test3", args{"abcd", "bc"}, []string{"a", "c"}},
	}
	for _, tt := range tests {

		//s.T().Run(tt.name, func(t *testing.T) {
		//	res := testcase.Split(tt.args.s, tt.args.sep)
		//	//assert.Equal(tt.wantResult, res, "not equal")
		//	require.Equal(s.T(), tt.wantResult, res, "not equal")
		//})
		s.Run(tt.name, func() {

			res := testcase.Split(tt.args.s, tt.args.sep)
			//assert.Equal(tt.wantResult, res, "not equal")
			assert.Equal(s.T(), tt.wantResult, res, "not equal")
		})
	}
}
