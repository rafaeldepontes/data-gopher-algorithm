# HashMap

The HashMap, or also known as dictionary or just map, is a very cool data structure that implements an associative array (basically a abstraction for key/pair collections).

The HashMap use a hash function to "create" a special index, also called hash code, into an array of buckets or slots.

A HashMap has a time complexity of O(1) for best and average cases and for the worst case O(n), and for the space complexity... is always O(n).

There are some concepts behind a hash map, like a hash function, collision and work load, but we will see that in the future! For now, let's just see how to use it.

## Golang example:

> Keep in mind that a hashmap never is guaranteed to be ordered... each interation over a hashmap can be different from the last one.

```go
package main

func main() {
	capacity := 3
	hashMap := make(map[int]string, capacity)

	hashMap[0] = "A"
	hashMap[1] = "B"
	hashMap[2] = "C"

	for index, value := range hashMap {
		println("Index:", index, "- Value", value)
	}
}
```

## 219. Contains Duplicate II

Given an integer array `nums` and an integer `k`, return `true` \*if there are two **distinct indices\*** `i` and `j` in the _array_ such that `nums[i] == nums[j]` and `abs(i - j) <= k`.

Example 1:

- Input: nums = [1,2,3,1], k = 3
- Output: true

Example 2:

- Input: nums = [1,0,1,1], k = 1
- Output: true

Example 3:

- Input: nums = [1,2,3,1,2,3], k = 2
- Output: false

## Code:

```go
package main

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func containsNearbyDuplicate(nums []int, k int) bool {
	hashMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if _, ok := hashMap[v]; !ok {
			hashMap[v] = i
		} else {
			if abs(i-hashMap[v]) <= k {
				return true
			} else {
				hashMap[v] = i
			}
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3
	println("Contains Nearby Duplicate?", containsNearbyDuplicate(nums, k))
}
```

