package leetcode

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	mergedArray := make([]int, length)
	if len(nums1) == 0 {
		mergedArray = nums2
	} else if len(nums2) == 0 {
		mergedArray = nums1
	} else {
		index1 := 0
		index2 := 0
		for index := 0; index < length/2+1; index++ {
			if index1 == len(nums1) {
				mergedArray[index] = nums2[index2]
				index2++
			} else if index2 == len(nums2) {
				mergedArray[index] = nums1[index1]
				index1++
			} else if nums1[index1] < nums2[index2] {
				mergedArray[index] = nums1[index1]
				index1++
			} else {
				mergedArray[index] = nums2[index2]
				index2++
			}
		}
	}
	if length%2 == 1 {
		return float64(mergedArray[length/2])
	} else {
		return float64(mergedArray[length/2]+mergedArray[length/2-1]) / 2.0
	}
}
