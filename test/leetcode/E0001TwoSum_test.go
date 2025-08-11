package leetcode

import (
	"github.com/techlandspace/algorithm-go/internal/leetcode"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	leetcode.TwoSum(nums, target)
}
