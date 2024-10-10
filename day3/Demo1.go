package day3

//获取字符串中最长连续不重复子串；如输入abcdabc，输出abcd

func getFirstNChars(s string, n int) string {
	var count int
	var result []rune

	for _, r := range s {
		result = append(result, r)
		count++
		if count == n {
			break
		}
	}

	return string(result)
}
