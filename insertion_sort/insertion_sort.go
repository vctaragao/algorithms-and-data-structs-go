package main

import "fmt"

func insertion_sort(data []int) {
	for i := 1; i < len(data); i++ {
		insert := data[i]
		j := i - 1

		for j >= 0 && insert < data[j] {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = insert
	}
}

func main() {
	sort([]int{6, 0, 3, 5})
	sort([]int{0, 1, 4, 2, 8})
	sort([]int{8, 6, 4, 2, 5, 9, 7, 3})
}

func sort(data []int) {
	fmt.Printf("unsorted: %v\n", data)
	insertion_sort(data)
	fmt.Printf("sorted: %v\n", data)
}
