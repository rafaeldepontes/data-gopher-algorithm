# Quick Sort

Quick-Sort is a very interesting sorting algorithm, it works by diveding the list in smaller sections and then sorting it! This type of algorithm is called "Divide and Conquer"

We choose the rightest number on our list as our pivot, then we go through the list checking if the number on our right is less or equal than our pivot, if it's we increase the `i`, which represents the starting of our list, and switch the `ith` number with the `jth` number! Then finally we put the pivot in the "right spot" on our list and then we simply return that pivot position! Now we do it again and again and again...

Time Complexity is O(n log n) and the Space Complexity is O(log n) for the best and average case. If happens to be the worst case possible then the time complexity becomes O(n^2) and the space complexity O(n)

## Go Implementation:

```go
package main

import "fmt"

// This returns the partition's position
func partition(nums []int, left, right int) int {
	pivot := nums[right]

	i := left - 1

	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[right] = nums[right], nums[i+1]

	return i + 1
}

func quickSort(nums []int, left, right int) {
	if left < right {
		pi := partition(nums, left, right)
		quickSort(nums, left, pi-1)
		quickSort(nums, pi+1, right)
	}
}

func main() {
	nums := []int{4, 12, 5, 72, 8, 17, 54, 26, 88, 92, 1}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println("Sorted nums:", nums)
}
```
