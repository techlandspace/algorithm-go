package leetcode

func subarraySum(nums []int, k int) int {
	var sum []int
    sum = append(sum, nums[0])
    count := 0
	if sum[0] == k {
		count++
	}
    for i := 1; i < len(nums); i++ {
        sum = append(sum, sum[i - 1] + nums[i])
		if(sum[i] == k){
			count++
		}
        for j := 0; j < i; j++ {
            if sum[i] - k == sum[j] {
                count++
            }
        }
    }
    return count
}

//上一个算法中显示存储了所有前缀和，因此导致了第二次for循环。
//实际上可以通过哈希表来存储所有前缀和，这样就可以直接查找是否存在前缀和为sum[i] - k的前缀和，从而避免了第二次for循环。
//同时，由于需要处理第一个元素的特殊情况，因此可以将哈希表的初始值设置为{0: 1}，这样就可以避免对第一个元素的特殊处理。
func subarraySum2(nums []int, k int) int {
	preSum := make(map[int]int)
	preSum[0] = 1 // 初始值为{0: 1}，避免对第一个元素的特殊处理
	count, sum := 0, 0
	for _, num := range nums {
		sum += num
		if preSum[sum - k] > 0 {
			count += preSum[sum - k] // 直接查找是否存在前缀和为sum[i] - k的前缀和
		}
		preSum[sum]++ // 将当前前缀和加入哈希表中
	}
	return count
}