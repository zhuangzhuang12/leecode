package day2

func isValid(s string) bool {
	queue := make([]rune, 0)
	for _, s1 := range s {
		if s1 == '{' || s1 == '[' || s1 == '(' {
			queue = append(queue, s1)
		} else if s1 == '}' {
			if len(queue) > 0 && queue[len(queue)-1] == '{' {
				queue = queue[:len(queue)-1]
			} else {
				return false
			}
		} else if s1 == ']' {
			if len(queue) > 0 && queue[len(queue)-1] == '[' {
				queue = queue[:len(queue)-1]
			} else {
				return false
			}
		} else if s1 == ')' {
			if len(queue) > 0 && queue[len(queue)-1] == '(' {
				queue = queue[:len(queue)-1]
			} else {
				return false
			}
		}
	}
	return len(queue) == 0
}
