package leetcode

import (
	"github.com/techlandspace/algorithm-go/internal/leetcode"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	t.Log(leetcode.FindMedianSortedArrays(nums1, nums2))
}
