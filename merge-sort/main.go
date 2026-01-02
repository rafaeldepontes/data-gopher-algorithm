package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func findMiddle(head *Node) *Node {
	slow, fast := head, head.next

	for fast != nil && fast.next != nil {
		fast = fast.next
		slow = slow.next
	}

	return slow
}

func merge(l1, l2 *Node) *Node {
	dummy := &Node{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.value < l2.value {
			curr.next = l1
			l1 = l1.next
		} else {
			curr.next = l2
			l2 = l2.next
		}
		curr = curr.next
	}

	if l1 != nil {
		curr.next = l1
	} else {
		curr.next = l2
	}

	return dummy.next
}

func mergeSort(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	middle := findMiddle(head)
	afterMiddle := middle.next
	middle.next = nil

	left := mergeSort(head)
	right := mergeSort(afterMiddle)

	return merge(left, right)
}

func main() {
	list := &Node{value: 9, next: &Node{value: 5, next: &Node{value: 1, next: &Node{value: 15, next: nil}}}}
	sortedList := mergeSort(list)
	formatedText := "%d -> "
	decimalText := "%d"

	for sortedList != nil {
		if sortedList.next == nil {
			fmt.Printf(decimalText, sortedList.value)
		} else {
			fmt.Printf(formatedText, sortedList.value)
		}
		sortedList = sortedList.next
	}
}
