package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 3, 6, 4, 3, 1, 6, 8, 5}
	fmt.Println(findRepeat(arr))
}

func findRepeat(s []int) []int {
	tmp := make(map[int]int)
	result := make([]int, 0)

	for _, v := range s {
		if _, ok := tmp[v]; !ok {
			tmp[v] = 1
		} else {
			tmp[v]++
		}
		if tmp[v] == 2 {
			result = append(result, v)
		}
	}
	return result
}
