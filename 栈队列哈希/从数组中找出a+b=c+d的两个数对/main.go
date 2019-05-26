package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(20))
	}
	fmt.Println(arr)
	fmt.Println(findFour(arr))
}

func findFour(arr []int) ([]int, error) {
	tmp := make(map[int][]int)
	length := len(arr)
	for i := 0; i < length; i++ {
		for k := i + 1; k < length; k++ {
			sum := arr[i] + arr[k]
			if v, ok := tmp[sum]; !ok {
				tmp[sum] = []int{arr[i], arr[k]}
			} else {
				return append(v, arr[i], arr[k]), nil
			}
		}
	}
	return nil, errors.New("none")
}
