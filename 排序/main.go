package main

import "fmt"

func main() {
	fmt.Println("排序算法")
	list := []int{16, 4, 3, 1, 9, 5, 1, 34, 7, 64, 24, 76, 987, 32}
	fmt.Println(list)
	kuaisu(list)
	fmt.Println(list)
}

func maopao(list []int) {
	l := len(list)
	for i := l - 1; i > 0; i-- {
		flag := false
		for k := 0; k < i; k++ {
			if list[k+1] < list[k] {
				list[k], list[k+1] = list[k+1], list[k]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func xuanze(list []int) {
	l := len(list)
	for i := 0; i < l; i++ {
		for k := i + 1; k < l; k++ {
			if list[k] < list[i] {
				list[k], list[i] = list[i], list[k]
			}
		}
	}
}

func charu(list []int) {
	l := len(list)
	for i := 1; i < l; i++ {
		for k := i; k > 0; k-- {
			if list[k] < list[k-1] {
				list[k], list[k-1] = list[k-1], list[k]
			}
		}
	}
}

func kuaisu(list []int) {
	quickSort(list, 0, len(list)-1)
}

func quickSort(list []int, l, r int) {
	if l < r {
		pivot := list[r]
		i := l
		for k := l; k < r; k++ {
			if list[k] < pivot {
				list[k], list[i] = list[i], list[k]
				i++
			}
		}
		list[r], list[i] = list[i], list[r]
		quickSort(list, l, i-1)
		quickSort(list, i+1, r)
	}
}
