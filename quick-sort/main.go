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
