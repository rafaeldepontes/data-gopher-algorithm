# Circular Linked List

Well, now that we are here I think the name is self explanatory... Is a linked list where the LAST prev node on our list is not null (or nil)... its actually our HEAD because the list is FINALLY circular... so every time that we interate over it we will see the same node at the same position as before!

There are two types of CLL(Circular Linked List), the singly implementation and the doubly implementation... I will do both, but they can easily be changed to adapt the needs.

## Go Implementatio of Singly Circular Linked List

```go
package main

type Node struct {
	Value int
	Next  *Node
}

type SinglyCircularLinkedList struct {
	Head   *Node
	Length int
}

func NewSinglyCLL(data int) *SinglyCircularLinkedList {
	node := &Node{Value: data}
	node.Next = node

	return &SinglyCircularLinkedList{
		Head:   node,
		Length: 1,
	}
}

func (list *SinglyCircularLinkedList) Append(data int) {
	newNode := &Node{Value: data}

	tail := list.Head
	for tail.Next != list.Head {
		tail = tail.Next
	}

	newNode.Next = list.Head
	tail.Next = newNode

	list.Length++
}

func main() {
	list := NewSinglyCLL(1)
	list.Append(2)

	current := list.Head
	length := list.Length * 2

	println("========================")
	for range length {
		println("Current Node:", current.Value)
		println("Next Node will be:", current.Next.Value)
		println("========================")
		current = current.Next
	}
}
```

## Go Implementation of Doubly Circular Linked List

```go
package main

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type DoublyCircularLinkedList struct {
	Head   *Node
	Length int
}

func NewDoublyCLL(data int) *DoublyCircularLinkedList {
	node := &Node{Value: data}
	node.Next = node
	node.Prev = node

	return &DoublyCircularLinkedList{
		Head:   node,
		Length: 1,
	}
}

func (list *DoublyCircularLinkedList) Append(data int) {
	newNode := &Node{Value: data}

	tail := list.Head.Prev

	newNode.Next = list.Head
	newNode.Prev = tail

	tail.Next = newNode
	list.Head.Prev = newNode

	list.Length++
}

func main() {
	list := NewDoublyCLL(1)
	list.Append(2)
	list.Append(3)

	current := list.Head
	iterations := list.Length * 2

	println("========================")
	for range iterations {
		println("Current Node:", current.Value)
		println("Next Node:", current.Next.Value)
		println("Previous Node:", current.Prev.Value)
		println("========================")
		current = current.Next
	}
}
```
