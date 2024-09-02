package day1

func SortArray(nums []int) []int {
	ptrNums := make([]*int, len(nums))
	for i := range nums {
		ptrNums[i] = &nums[i] // 正确地指向 nums 中的元素
	}

	quick(ptrNums, 0, len(ptrNums)-1)
	return nums // 直接返回 nums
}

func quick(nums []*int, left, right int) {
	if left < right {
		index := getIndex(nums, left, right)
		quick(nums, left, index-1)
		quick(nums, index+1, right)
	}
}

func getIndex(nums []*int, left, right int) int {
	t := *nums[left] // 解引用 t
	for left < right {
		for left < right && *nums[right] >= t { // 解引用 nums[right]
			right--
		}
		for left < right && *nums[left] <= t { // 解引用 nums[left]
			left++
		}
		nums[left], nums[right] = nums[right], nums[left] // 交换指针
	}
	return left
}
