package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 先排序
	res := [][]int{}
	n := len(nums)

	for i := 0; i < n-2; i++ {
		// 跳过重复值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			switch {
			case sum == 0:
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 跳过重复值
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			case sum < 0:
				left++
			default:
				right--
			}
		}
	}
	return res
}
