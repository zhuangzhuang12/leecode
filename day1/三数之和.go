package day1

import "slices"

//	func ThreeSum(nums []int) [][]int {
//		res := [][]int{}
//		slices.Sort(nums)
//		var l1, l2, l3 int
//		for l1 = 0; l1 < len(nums)-2; l1++ {
//			if l1 > 0 && nums[l1] == nums[l1-1] {
//				continue
//			}
//			l2, l3 = l1+1, len(nums)-1
//			for l2 < l3 {
//				sum := nums[l1] + nums[l2] + nums[l3]
//				if sum == 0 {
//					res = append(res, []int{nums[l1], nums[l2], nums[l3]})
//					if l2 < l3 && nums[l2] == nums[l2+1] {
//						l2++
//					}
//					if l3 > l2 && nums[l3] == nums[l3-1] {
//						l3--
//					}
//					l2++
//					l3--
//				} else if sum < 0 {
//					l2++
//				} else {
//					l3--
//				}
//			}
//
//		}
//		return res
//	}
func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	slices.Sort(nums)
	var l1, l2, l3 int
	for l1 = 0; l1 < len(nums)-2; l1++ {
		if l1 > 0 && nums[l1] == nums[l1-1] {
			continue
		}
		l2, l3 = l1+1, len(nums)-1
		for l2 < l3 {
			if nums[l1]+nums[l2]+nums[l3] == 0 {
				res = append(res, []int{nums[l1], nums[l2], nums[l3]})
				if l2 < l3 && nums[l2] == nums[l2+1] {
					l2++
				}
				if l2 < l3 && nums[l3] == nums[l3-1] {
					l3--
				}
				l2++
				l3--
			} else if nums[l1]+nums[l2]+nums[l3] < 0 {
				l2++
			} else if nums[l1]+nums[l2]+nums[l3] > 0 {
				l3--
			}
		}
	}
	return res
}
