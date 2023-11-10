package richestcustomerwealth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (s *testSuite) TestGeneral1() {
	type test struct {
		name           string
		givenAccounts  [][]int
		expectedAnswer int
	}
	tests := []test{
		{
			name: "1",
			givenAccounts: [][]int{
				{1, 2, 3},
				{3, 2, 1},
			},
			expectedAnswer: 6,
		},
		{
			name: "2",
			givenAccounts: [][]int{
				{1, 5},
				{7, 3},
				{3, 5},
			},
			expectedAnswer: 10,
		},
		{
			name: "3",
			givenAccounts: [][]int{
				{2, 8, 7},
				{7, 1, 3},
				{1, 9, 5},
			},
			expectedAnswer: 17,
		},
	}
	for _, test := range tests {
		got := maximumWealth(test.givenAccounts)
		s.Equal(test.expectedAnswer, got, fmt.Sprintf("test %q not passed", test.name))
	}
}

type testSuite struct {
	suite.Suite
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, new(testSuite))
}
