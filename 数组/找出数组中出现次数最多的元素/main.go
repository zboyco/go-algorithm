package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 3, 6, 4, 3, 1, 6, 8, 5}
	fmt.Println(findMaxRepeat(arr))
}

func findMaxRepeat(s []int) int {
	tmp := make(map[int]int)
	result := 0
	maxCount := 0

	for _, v := range s {
		if _, ok := tmp[v]; !ok {
			tmp[v] = 1
		} else {
			tmp[v]++
		}
		if tmp[v] > maxCount {
			result = v
			maxCount = tmp[v]
		}
	}
	return result
}
