package day1

func longestConsecutive(nums []int) int {
	mmap := make(map[int]bool)
	for _, num := range nums {
		mmap[num] = true
	}
	res := 0
	for _, num := range nums {
		if mmap[num-1] {
			continue
		} else {
			cur := num
			curLen := 1
			for mmap[cur+1] {
				cur++
				curLen++
			}
			res = max(res, curLen)
		}
	}
	return res
}
