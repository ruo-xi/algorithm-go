package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(twoSum(nums, 5))
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, v := range nums {
		if index, ok := numMap[target-v]; ok {
			return []int{index, i}
		}

		numMap[v] = i
	}

	//	for i, v := range nums {
	//		index, ok := numMap[target-v]
	//		if ok {
	//			return []int{index, i}
	//		} else {
	//			numMap[v] = i
	//		}
	//	}

	return []int{0, 0}
}

