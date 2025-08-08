package leetcode

func maxArea(height []int) int {
    left, right := 0, len(height) - 1
    maxArea := 0
    for left < right {
        area := 0
        if height[left] < height[right] {
            area = (right - left) * height[left]
            left++
        } else {
            area = (right - left) * height[right]
            right--
        }
        if area > maxArea {
            maxArea = area
        }
    }
    return maxArea
}
