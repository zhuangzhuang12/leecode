package day1

import "slices"

//	func lengthOfLongestSubstring(s string) int {
//		res := 0
//		chacaterMap := make(map[string]int)
//		var i, j int
//		for i < len(s) {
//			c := s[i]
//			chacaterMap[string(c)]++
//			if chacaterMap[string(c)] > 1 {
//				for j < i {
//					if chacaterMap[string(c)] > 1 {
//						chacaterMap[string(s[j])]--
//						j++
//					} else {
//						break
//					}
//				}
//			}
//			i++
//			res = max(res, i-j)
//		}
//		return res
//	}
func threeSum(nums []int) [][]int {
	res := [][]int{}
	slices.Sort(nums)
	var l1, l2, l3 int
	for l1 = 0; l1 < len(nums)-2; l1++ {
		if l1 > 0 && nums[l1] == nums[l1-1] {
			continue
		}
		l3, l3 = l1+1, len(nums)-1
		for l2 < l3 {
			sum := nums[l1] + nums[l2] + nums[l3]
			if sum == 0 {
				res = append(res, []int{nums[l1], nums[l2], nums[l3]})
				if l2 < l3 && nums[l2] == nums[l2+1] {
					l2++
				}
				if l3 > l2 && nums[l3] == nums[l3-1] {
					l3--
				}
				l2++
				l3--
			} else if sum < 0 {
				l2++
			} else {
				l3--
			}
		}

	}
	return res
}
