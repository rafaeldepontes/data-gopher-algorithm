# Sliding Window

Sliding window is generally used to solve problems that involve finding, manipulating, or adding something new to sub-arrays(or vectors) or sub-strings (technically they are the same thing, but let's ignore that for now...) so that a specific action is performed or a result occurs.

## Golang Implementation:

This LeetCode problem can be easily solved with a sliding window approach:

### 3090. Maximum Length Substring With Two Occurrences

Given a string `s`, return the **maximum** length of a substring such that it contains at _most_ two occurrences of each character.

Example 1:

- Input: s = "bcbbbcba"
- Output: 4
- Explanation:
- The following substring has a length of 4 and contains at most two occurrences of each character: "bcbbbcba".

### Code:

```go
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
```
