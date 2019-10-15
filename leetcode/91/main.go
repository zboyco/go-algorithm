package main

import "fmt"

func main() {
	fmt.Println(numDecodings("226"))
}

// 动态规划
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	prev := 1
	curr := 1
	for i := 1; i < len(s); i++ {
		tmp := curr
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				curr = prev
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			curr += prev
		}
		prev = tmp
	}

	return curr
}

// 递归
func numDecodings1(s string) int {
	return decode(s, len(s)-1)
}

func decode(chars string, index int) int {
	if index <= 0 {
		if chars[0] == '0' {
			return 0
		}
		return 1
	}

	result := 0
	curr := chars[index]
	prev := chars[index-1]

	if curr > '0' {
		result = decode(chars, index-1)
	}

	if prev == '1' || (prev == '2' && curr <= '6') {
		result += decode(chars, index-2)
	}

	return result
}
