package main

import (
	"fmt"
	"strconv"
)

func isPalindRome(num int) bool {
	numStr := strconv.Itoa(num)

	for i, j := 0, len(numStr)-1; i < j; i, j = i+1, j-1 {
		if numStr[i] != numStr[j] {
			return false
		}
	}
	return true

}

func main() {
	num := 1234321
	result := isPalindRome(num)
	if result {
		fmt.Printf("%d 是回文!\n", num)
	} else {
		fmt.Printf("%d 不是回文!\n", num)
	}

}
