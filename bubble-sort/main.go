package main

import "fmt"

func bubbleSort(input []int) {
	n := len(input)
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}

		n--
	}
}

func main() {
	edgeCases := make(map[int][]int)
	case1 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	case2 := []int{4, 4, 6, 1}
	case3 := []int{}
	edgeCases[0] = case1
	edgeCases[1] = case2
	edgeCases[2] = case3

	for i := range len(edgeCases) {
		bubbleSort(edgeCases[i])
		fmt.Println("Case number:", i+1, "- Sorted list:", edgeCases[i])
	}
}
