package leetcode

// MoveZeroes moves all 0's to the end of the array while maintaining the relative order of the non-zero elements.

func moveZeroes(nums []int) {
	// Initialize a pointer to keep track of the position to place the next non-zero element.
	zeroCounts := 0
	// Iterate through the array.
	for i := 0; i < len(nums); i++ {
		// If the current element is non-zero, move it to the position pointed by nonZeroPointer.
		if nums[i] == 0 {
			zeroCounts++
		} else if zeroCounts > 0 {
			nums[i-zeroCounts] = nums[i]
			nums[i] = 0
		}
	}
}
