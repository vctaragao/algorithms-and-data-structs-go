package main

import "fmt"

func bubble_sort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func main() {
	sort([]int{6, 0, 3, 5})
	sort([]int{5, 1, 4, 2, 8})
	sort([]int{8, 6, 4, 2, 5, 9, 7, 3})
}

func sort(data []int) {
	fmt.Printf("unsorted: %v\n", data)
	bubble_sort(data)
	fmt.Printf("sorted: %v\n", data)
}
