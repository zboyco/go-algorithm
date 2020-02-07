package main

import "fmt"

func main() {
	sss := "nnabac"

	fmt.Println(longestPalindrome(sss))
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		maxLen := center(s, i, i)
		len2 := center(s, i, i+1)
		if maxLen < len2 {
			maxLen = len2
		}
		if maxLen > (end - start + 1) {
			start = i - (maxLen-1)/2
			end = start + maxLen - 1
		}
	}
	return s[start : end+1]
}

func center(s string, left, right int) int {
	maxIndex := len(s) - 1
	for left >= 0 && right <= maxIndex && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
