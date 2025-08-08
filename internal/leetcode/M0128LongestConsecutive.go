package leetcode

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	sort.Ints(nums)
	max := 1
	current := 1
	for i := 0; i < len(nums); i++ {
		if i == len(nums) - 1 {
			if current > max {
				max = current
			}
			break
		}
		if nums[i] == nums[i+1]  {
			continue
		} else if nums[i] == nums[i+1]-1 {
			current++
		} else {
			if current > max {
				max = current
			}
			current = 1
		}
	}
	return max
}