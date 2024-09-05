package day1

func twoSum(nums []int, target int) []int {
	nmap := make(map[int]int)
	for i, num := range nums {
		if j, ok := nmap[target-num]; ok {
			return []int{j, i}
		} else {
			nmap[num] = i
		}
	}
	return []int{}
}
