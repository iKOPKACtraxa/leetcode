package leetcode

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func (s *SolvingSuit) Test1() {
	get := runningSum([]int{1, 2, 3, 4})
	s.Equal([]int{1, 3, 6, 10}, get)
}

func (s *SolvingSuit) Test2() {
	get := runningSum([]int{1, 1, 1, 1, 1})
	s.Equal([]int{1, 2, 3, 4, 5}, get)
}

func (s *SolvingSuit) Test3() {
	get := runningSum([]int{3, 1, 2, 10, 1})
	s.Equal([]int{3, 4, 6, 16, 17}, get)
}

type SolvingSuit struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SolvingSuit))
}
