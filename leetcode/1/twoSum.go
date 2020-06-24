package main

import "fmt"

func TwoSumBruteForce(nums []int, target int) [2]int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == target-nums[j] {
				return [2]int{i, j}
			}
		}
	}
	panic("No Two sum solution")
}

func TwoSumTwoPassHashTable(nums []int, target int) [2]int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if v, ok := m[complement]; ok && i != v {
			return [2]int{i, v}
		}
	}
	panic("No two sum solution")
}

func TwoSumOnePassHashTable(nums []int, target int) [2]int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if v, ok := m[complement]; ok {
			return [2]int{v, i}
		} else {
			m[nums[i]] = i
		}
	}
	panic("No two sum solution")
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	target := 3
	fmt.Println(TwoSumBruteForce(nums, target))
	fmt.Println(TwoSumTwoPassHashTable(nums, target))
	fmt.Println(TwoSumOnePassHashTable(nums, target))
}
