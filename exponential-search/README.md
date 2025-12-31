# Exponential Search

Exponential search uses binary search after a certain condition is met. It searches for the target by checking if it is greater or less than a number X. If it is greater than X, the number X will be doubled, if it is less... Binary search will be applied to the only part of the list that actually matters. It has the same limitation as binary search, therefore, the list needs to be sorted.

## Go Implementation

```go
package main

import (
	"errors"
	"math"
)

func binarySearch(nums []int, target, head, tail int) (int, error) {
	var middle int

	for head < tail {
		middle = (head - (head-tail)/2)
		if nums[middle] < target {
			head = middle + 1
		} else if nums[middle] > target {
			tail = middle - 1
		} else if nums[middle] == target {
			return middle, nil
		}
	}
	return 0, errors.New("Target is not on the list")
}

func exponentialSearch(nums []int, target int) (int, error) {
	if target == 0 {
		return 0, errors.New("Target cannot be 0")
	}

	if nums[0] == target {
		return 0, nil
	}

	var (
		pos    = 1
		length = len(nums) - 1
	)

	for pos < length && nums[pos] <= target {
		pos *= 2
	}

	return binarySearch(nums, target, (pos / 2), int(math.Min(float64(pos), float64(length-1))))
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	target := 4
	result, err := exponentialSearch(nums, target)
	if err != nil {
		println(err.Error())
		return
	}
	println("Result:", result)
}
```
