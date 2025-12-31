# Two Pointers

It's basically the initialization of the variables that needs to hold different memory spaces, each one go through an array (normally O(n) time complexity) and do some tipe of manipulation in it, it can be reversing it, duplicating or putting it in some specific order...

## Golang example of a Two Points algorithm usage:

- This leet code problem can be handle easily with a two pointer approach:

```md
2000. Reverse Prefix of Word

Given a 0-indexed string word and a character ch, reverse the segment of word that starts at index 0 and ends at the index of the first occurrence of ch (inclusive). If the character ch does not exist in word, do nothing.

- For example, if word = "abcdefd" and ch = "d", then you should reverse the segment that starts at 0 and ends at 3 (inclusive). The resulting string will be "dcbaefd".

Return the resulting string.

Example 1:

    Input: word = "abcdefd", ch = "d"
    Output: "dcbaefd"
    Explanation: The first occurrence of "d" is at index 3.
    Reverse the part of word from 0 to 3 (inclusive), the resulting string is "dcbaefd".
```

Code:

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
