package day1

import "slices"

func LongestPalindrome(s string) int {
	nummap := map[string]int{}
	for i := 0; i < len(s); i++ {
		nummap[string(s[i])]++
	}
	jishu := []int{}
	oushu := []int{}
	for _, num := range nummap {
		if num%2 != 0 {
			jishu = append(jishu, num)
		} else {
			oushu = append(oushu, num)
		}
	}
	slices.Sort(jishu)
	res := 0
	for _, num := range oushu {
		res = res + num
	}
	if len(jishu) >= 1 {
		res += jishu[len(jishu)-1]
	}

	return res
}
