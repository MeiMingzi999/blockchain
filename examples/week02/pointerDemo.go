package main

import "fmt"

func updateInt(pointer *int) int {
	return *pointer + 10
}

func updateSlice(numList *[]int) []int {
	nums := *numList
	for index, num := range nums {
		num = num * 2
		nums[index] = num
	}
	return nums
}

func main() {
	test := 10
	fmt.Printf("修改后的值为: %d\n", updateInt(&test))

	testList := []int{1, 2, 3, 4}
	fmt.Printf("修改后的切片为: %v\n", updateSlice(&testList))

}
