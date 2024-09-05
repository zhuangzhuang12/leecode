package day1

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (r + l) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[l] {
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target <= nums[r] && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
