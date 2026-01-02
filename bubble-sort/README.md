# Bubble Sort

One of the most known sorting algorithms in the world, the Bubble Sort is a classic and should be treat with respect! It's a algorithm really slow, but it has their uses.

Bubble sort is really simples, you just need to do a for loop inside of a for loop and compare if the number `ith` that you're currently looking at is higher or lower than the adjacent `ith`(or the `i + 1`...), if it's higher then you need to move the number `ith` to the `ith+1` position.

You need to keep doing that until you finally sort the list.

If none is swapped... well... the algorithm should do a earlier break.

## Go Implementatio:

```go
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
```
