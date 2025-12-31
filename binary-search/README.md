# Binary Search

This algorithm is really interesting because is done in place, so it doesn't create a new data structure for it search! It only use two up to three variables to help it locate itself in the array (for example.)

Binary Search has a time complexity of O(logn) and a space complexity of O(1)

The only problem with binary search, is that this algorithm only works on ordered lists.

## Golang implementation:

```go
package main

import "errors"

func binarySearch(nums []int, target int) (int, error) {
	var (
		middle int
		left   int = 0
		right  int = len(nums) - 1
	)

	for left < right {
		middle = (left - (left-right)/2)
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] == target {
			return middle, nil
		}
	}
	return 0, errors.New("Target is not on the list")
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	target := 21
	ans, err := binarySearch(nums, target)
	if err != nil {
		println("[ERROR]", err.Error())
		return
	}
	println(ans)
}
```
