package main

import "fmt"

func longestCommonPrefix(s []string) string {
	prefixMap := make([]byte, 0)
	firstStr := s[0]

	for i := 0; i < len(firstStr); i++ {
		for index, str := range s {
			if index == 0 {
				continue
			}
			if len(str) < i+1 {
				return string(prefixMap)
			}
			if firstStr[i] != str[i] {
				return string(prefixMap)
			}
		}
		prefixMap = append(prefixMap, firstStr[i])

	}
	return string(prefixMap)

}

func main() {
	test := []string{"ababababab", "abab", "a"}
	fmt.Println(longestCommonPrefix(test))
}
