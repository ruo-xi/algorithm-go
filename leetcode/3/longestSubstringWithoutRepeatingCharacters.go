package main

type void struct {
}

var member void

func lengthOfLongestSubstringBruteForce(s string) int {
	cs := []rune(s)
	n := len(cs)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if allUnique(cs, i, j) && j-i+1 > ans {
				ans = j - i + 1
			}
		}
	}
	return ans
}

func lengthOfLongestSubstringSlideWindow(s string) int {
	cs, set, ans, i, j := []rune(s), make(map[rune]void), 0, 0, 0
	n := len(cs)
	for i < n && j < n {
		if _, ok := set[cs[j]]; ok {
			delete(set, cs[i])
			i++
		} else {
			set[cs[j]] = member
			j++
			if ans < j-i {
				ans = j - i
			}
		}
	}
	return ans
}

func lengthOfLongestSubstringSlideWindowOptimized(s string) int {
	cs := []rune(s)
	n, ans := len(cs), 0
	m := make(map[rune]int)
	for i, j := 0, 0; j < n; j++ {
		if _, ok := m[cs[j]]; ok {
			if m[cs[j]] > i {
				i = m[cs[j]]
			}
		}
		if ans < j-i+1 {
			ans = j - i + 1
		}
		m[cs[j]] = j + 1
	}
	return ans
}
func main() {

}

func allUnique(s []rune, start, end int) bool {
	m := make(map[rune]void)
	for i := start; i <= end; i++ {
		if _, ok := m[s[i]]; ok {
			return false
		}
		m[s[i]] = member
	}
	return true
}
