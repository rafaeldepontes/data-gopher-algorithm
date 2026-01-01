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
	capacity := 3
	hashMap := make(map[int]string, capacity)

	hashMap[0] = "A"
	hashMap[1] = "B"
	hashMap[2] = "C"

	for index, value := range hashMap {
		println("Index:", index, "- Value", value)
	}

	println()

	nums := []int{1, 2, 3, 1}
	k := 3
	println("Contains Nearby Duplicate?", containsNearbyDuplicate(nums, k))
}
