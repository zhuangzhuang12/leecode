package day1

//func lengthOfLongestSubstring(s string) int {
//	res := 0
//	cmap := map[string]int{}
//	j := 0
//	for i := 0; i < len(s); i++ {
//		c := s[i]
//		cmap[string(c)]++
//		if cmap[string(c)] > 1 {
//			for j < i {
//				if cmap[string(c)] > 1 {
//					cmap[string(s[j])]--
//					j++
//				} else {
//					break
//				}
//			}
//		}
//		res = max(res, (i - j + 1))
//	}
//	return res
//}

func lengthOfLongestSubstring(s string) int {
	res := 0
	ma := map[string]int{}
	j := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		ma[string(c)]++

		for ; ma[string(c)] > 1 && j < i; j++ {
			a := s[j]
			ma[string(a)]--
		}
		res = max(res, i-j+1)
	}

	return res

}
