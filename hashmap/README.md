# HashMap

The HashMap, or also known as dictionary or just map, is a very cool data structure that implements an associative array (basically a abstraction for key/pair collections).

The HashMap use a hash function to "create" a special index, also called hash code, into an array of buckets or slots.

A HashMap has a time complexity of O(1) for best and average cases and for the worst case O(n), and for the space complexity... is always O(n).

There are some concepts behind a hash map, like a hash function, collision and work load, but we will se that in the future! For now, let's just see how to use.

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
