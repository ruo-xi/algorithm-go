package main

import (
	"fmt"
)

func main() {
	s := "abcdabca"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {

	/* 不那么节省空间 */
	runes := []rune(s)
	var result, start int = 0, -1
	runeMap := make(map[rune]int)

	for i, v := range runes {
		if index, ok := runeMap[v]; ok {
			start = max(index, start)
		}
		result = max(i-start, result)
		runeMap[v] = i
	}

	/* 稍微节省空间  */
	// if len(s) == 0 {
	// 	return 1
	// }
	// runes := []rune(s)
	// var sum, start, result int
	// runeMap := make(map[rune]int)
	// for i, v := range runes {
	// 	if index, ok := runeMap[v]; ok {
	// 		if sum > result {
	// 			result = sum
	// 		}
	// 		for start <= index {
	// 			delete(runeMap, runes[start])
	// 			start++
	// 			sum--
	// 		}
	// 	}
	// 	sum++
	// 	runeMap[v] = i
	// }
	// if sum > result {
	// 	result = sum
	// }

	return result
}

func max(x int, y int) int {
	if x < y {
		return y
	}

	return x
}
