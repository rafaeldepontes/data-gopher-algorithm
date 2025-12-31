package main

import "sort"

func countPairs(nums []int, target int) int {
	sort.Ints(nums) // This makes the time complexity O(nlogn)
	var (
		ans   int
		left  int = 0
		right int = len(nums) - 1
	)

	for left < right {
		if nums[left]+nums[right] < target {
			ans += right - left
			left++
		} else {
			right--
		}
	}
	return ans
}
