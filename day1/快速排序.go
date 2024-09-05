package day1

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	index := len(nums) / 2
	left := []int{}
	mid := []int{}
	right := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[index] {
			left = append(left, nums[i])
		} else if nums[i] > nums[index] {
			right = append(right, nums[i])
		} else {
			mid = append(mid, nums[i])
		}
	}
	left = sortArray(left)
	right = sortArray(right)
	return append(append(left, mid...), right...)
}
