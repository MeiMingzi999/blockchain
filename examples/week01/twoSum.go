package main

import (
	"fmt"
)

// 两数之和, 返回两数下标
func twoSum(nums []int, target int) []int {
	result := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = []int{i, j}
			}
		}
	}
	return result

}

func main() {
	test := []int{1, 2, 3, 4}
	fmt.Print(twoSum(test, 3))
}
