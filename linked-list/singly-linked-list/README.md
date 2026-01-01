# Singly Linked List or SLL

A singly linked list is the simplest kind of linked lists. It takes up less space in memory because each node has only one address to the next node. So basically one value and one next.

## Golang Implementation

```go
package main

type Node struct {
	value int
	next  *Node
}

func NewLinkedList(data int) *Node {
	return &Node{
		value: data,
		next:  nil,
	}
}

func (l Node) NewNode(data int) *Node {
	return &Node{
		value: data,
		next:  nil,
	}
}

func main() {
	sll := NewLinkedList(1)
	sll.next = sll.NewNode(2)

	for sll != nil {
		println(sll.value)
		sll = sll.next
	}
}
```
