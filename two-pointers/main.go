package main

import "fmt"

func main() {
	input := "abcdefd"
	var prefix byte = 'd'
	fmt.Println("Input:", input, "Prefix:", prefix)
	fmt.Println("Reversed prefix:", reversePrefix(input, prefix), "\n")

	nums := []int{-1, 1, 2, 3, 1}
	target := 2
	fmt.Println("Nums:", nums, "Target", target)
	fmt.Println("Count Pairs:", countPairs(nums, target))
}
