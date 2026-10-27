package main

import "fmt"

// 找到出现一次的数字
func findOneNumber(numberList []int) []int {
	numberCount := make(map[int]int)
	result := make([]int, 0)
	// key 为数字, value为次数
	for _, value := range numberList {
		numberCount[value]++
	}
	if len(numberCount) > 0 {
		for key, count := range numberCount {
			if count == 1 {
				result = append(result, key)
			}
		}
	}
	return result

}

func main() {
	testList := []int{1, 1, 1, 12, 2, 2, 2, 3}
	result := findOneNumber(testList)
	fmt.Println(result)
}
