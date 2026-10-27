package main

import (
	"fmt"
	"strconv"
)

// 加一
func plusOne(digits []int) []int {
	numStr := ""
	numList := make([]int, 0)
	for _, num := range digits {
		numStr += strconv.Itoa(num)
	}
	fmt.Printf("numStr = %v\n", numStr)
	numInt, err := strconv.Atoi(numStr)
	if err == nil {
		numInt = numInt + 1
		numStr = strconv.Itoa(numInt)
		for _, str := range numStr {
			numList = append(numList, int(str-'0'))
		}
	}
	return numList
}

func main() {
	test := []int{1, 2, 3}
	fmt.Println(plusOne(test))
}
