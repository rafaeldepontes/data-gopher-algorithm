package main

import (
	"fmt"
	"math"
)

func maximumLengthSubstring(s string) int {
	var (
		max   int = 1
		left  int = 0
		right int = 0
	)

	count := make([]int, 26)
	count[s[right]-'a'] = 1

	for right < len(s)-1 {
		right++
		count[s[right]-'a'] += 1

		for count[s[right]-'a'] >= 3 {
			count[s[left]-'a'] -= 1
			left++
		}

		max = int(math.Max(float64(max), float64((right-left)+1)))
	}

	return max
}

func main() {
	s := "xyzzaz"
	fmt.Println("Good Substrings", maximumLengthSubstring(s))
}
