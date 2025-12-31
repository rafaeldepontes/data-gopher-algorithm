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
