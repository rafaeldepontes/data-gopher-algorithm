# Doubly Linked List

The SLL (Singly Linked List) are the base for the doubly implementation, here we have a new node called previous (or prev) that reference to the previous Node, different from the SLL the DLL (Doubly Linked List) now has a bi-directional structure allowing us to go back and forth on our list.

## Go Implementation

```go
package main

type DoublyLinkedList struct {
	Value int
	Next  *DoublyLinkedList
	Prev  *DoublyLinkedList
}

func NewLinkedList(data int) *DoublyLinkedList {
	return &DoublyLinkedList{
		Value: data,
		Next:  nil,
		Prev:  nil,
	}
}

func (list *DoublyLinkedList) NewNode(data int, prev *DoublyLinkedList) *DoublyLinkedList {
	return &DoublyLinkedList{
		Value: data,
		Next:  nil,
		Prev:  prev,
	}
}

func main() {
	dll := NewLinkedList(1)
	dll.Next = dll.NewNode(2, dll)

	println("======================")
	for dll != nil {
		println("Current Node:", dll.Value)
		println("It has a previous node?", dll.Prev != nil)
		if dll.Prev != nil {
			println("Previous Node:", dll.Prev.Value)
		}
		println("======================")
		dll = dll.Next
	}
}
```
