package main

import "fmt"

func main() {
	fmt.Println("239")

	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}

	fmt.Println(maxSlidingWindow(nums, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := make([]int, len(nums)-k+1)

	for i := range nums {
		if i > k - 1 {

		}
		
	}
	return result
}
