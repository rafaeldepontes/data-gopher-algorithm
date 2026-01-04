# Two Pointers

It's basically the initialization of two variables that needs to hold different memory spaces, each one go through an array (normally O(n) time complexity) and do some type of manipulation in it, it can be reversing it, duplicating or putting it in some specific order...

## Golang example of a Two Points algorithm usage:

- This LeetCode problem can be easily solved with a two pointer approach:
  <u>This text will be underlined.</u>

### 2000. Reverse Prefix of Word

Given a **0-indexed** string `word` and a character `ch`, **reverse** the segment of `word` that starts at index `0` and ends at the index of the **first occurrence** of `ch` (**inclusive**). If the character `ch` does not exist in `word`, do nothing.

- For example, if `word = "abcdefd"` and `ch = "d"`, then you should **reverse** the segment that starts at `0` and ends at `3` (**inclusive**). The resulting string will be `"dcbaefd"`.

Return the _resulting_ string.

Example 1:

- Input: word = "abcdefd", ch = "d"
- Output: "dcbaefd"
- Explanation: The first occurrence of "d" is at index 3.
- Reverse the part of word from 0 to 3 (inclusive), the resulting string is "dcbaefd".

### Code:

```go
package main

import "strings"

func reverseString(word []byte) string {
	left, right := 0, len(word)-1
	for left <= right {
		word[left], word[right] = word[right], word[left]
		left++
		right--
	}
	return string(word)
}

func reversePrefix(word string, ch byte) string {
	ans := strings.Builder{}
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			ans.WriteString(reverseString([]byte(word[:i+1])))
			ans.WriteString(word[i+1:]) // once we reverse the original prefix, we join the rest.
			break
		}
	}

    // If the prefix isn't in the string, then you return the original input
	if ans.Len() == 0 {
		return word
	}

	return ans.String()
}

func main() {
	input := "abcdefd"
	var prefix byte = 'd'
	println(reversePrefix(input, prefix))
}
```

And this problem also can be solved with a two-pointers approach:

## 2824. Count Pairs Whose Sum is Less than Target

Given a **0-indexed** integer array `nums` of length `n` and an integer `target`, return the number of pairs `(i, j)` where `0 <= i < j < n` and `nums[i] + nums[j] < target`.

Example 1:

- Input: nums = [-1,1,2,3,1], target = 2
- Output: 3
- Explanation: There are 3 pairs of indices that satisfy the conditions in the statement:

  - (0, 1) since 0 < 1 and nums[0] + nums[1] = 0 < target
  - (0, 2) since 0 < 2 and nums[0] + nums[2] = 1 < target
  - (0, 4) since 0 < 4 and nums[0] + nums[4] = 0 < target

- Note that (0, 3) is not counted since nums[0] + nums[3] is not strictly less than the target.

### Code:

```go
package main

import (
	"fmt"
	"sort"
)

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

func main() {
	nums := []int{-1, 1, 2, 3, 1}
	target := 2
	fmt.Println("Nums:", nums, "Target", target)
	fmt.Println("Count Pairs:", countPairs(nums, target))
}
```

