package leetcode

// 3. Longest Substring Without Repeating Characters
// func lengthOfLongestSubstring(s string) int {
// 	if len(s) == 0 {
// 		return 0
// 	}
// 	seen := make(map[byte]int)
// 	maxLength := 0
// 	start := 0
// 	for i, value := range s {
// 		if index, ok := seen[byte(value)]; ok && index >= start {
// 			start = index + 1
// 		}
// 		seen[byte(value)] = i
// 		if i-start+1 > maxLength {
// 			maxLength = i - start + 1
// 		}
// 	}
// 	return maxLength
// }

func lengthOfLongestSubstring(s string) int {
	left := 0
	chars := make(map[byte]int)
	maxLen := 0
	for index, str := range s {
		if value, ok := chars[byte(str)]; !ok || value < left {
			if maxLen < index-left+1 {
				maxLen = index - left + 1
			}
		} else {
			left = chars[byte(str)] + 1
		}

		chars[byte(str)] = index
	}
	return maxLen
}
