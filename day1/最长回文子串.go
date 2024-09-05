package day1

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res1 := maxString(s, i, i)
		res2 := maxString(s, i, i+1)
		if len(res1) > len(res2) && len(res1) > len(res) {
			res = res1
		} else if len(res2) > len(res) {
			res = res2
		}
	}
	return res
}

func maxString(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return s[i+1 : j]
}
