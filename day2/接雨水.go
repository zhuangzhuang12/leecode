package day2

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	var (
		left  []int
		right []int
	)
	left = make([]int, len(height))
	right = make([]int, len(height))
	left[0] = height[0]
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}
	right[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}
	var sum int
	for i := 0; i < len(height); i++ {
		sum += min(left[i], right[i]) - height[i]
	}
	return sum
}
