package day1

func SubarraySum(nums []int, k int) int {
	preSum := map[int]int{}
	preSum[0] = 1
	res := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if num, ok := preSum[sum-k]; ok {
			res += num
		}
		preSum[sum]++
	}
	return res
}
