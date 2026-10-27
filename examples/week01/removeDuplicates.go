package main

import "fmt"

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

func main() {
	nums := []int{1, 1, 1, 2, 3, 4}
	k := removeDuplicates(nums)
	fmt.Printf("长度是: %d, 数组前k个: %v\n", k, nums[:k])
}
