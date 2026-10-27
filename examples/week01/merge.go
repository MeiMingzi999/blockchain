package main

import (
	"fmt"
	"sort"
)

// 合并区间
func merge(intervals [][]int) [][]int {
	//先根据第一维排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	templt := make([][]int, 0)

	for i := 0; i < len(intervals); i++ {
		left := intervals[i][0]
		right := intervals[i][1]
		for i+1 < len(intervals) && right >= intervals[i+1][0] {
			right = max(right, intervals[i+1][1])
			i++
		}
		templt = append(templt, []int{left, right})
	}
	return templt

}

func main() {
	test := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(test)) // [[1 6] [8 10] [15 18]]
	test2 := [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(test2)) // [[1 5]]
}
