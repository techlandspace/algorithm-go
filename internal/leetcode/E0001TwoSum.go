package leetcode

import "fmt"

func TwoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if j, exists := numsMap[target-num]; exists {
			fmt.Println([]int{j, i})
			return []int{j, i}
		}
		numsMap[num] = i
	}
	return nil
}
