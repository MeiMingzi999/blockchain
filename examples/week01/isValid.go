package main

// 有效括号
func isValid(s string) bool {
	stringMap := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	stack := make([]byte, 0)

	if len(s)%2 == 1 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if _, ok := stringMap[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != stringMap[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return true

}
