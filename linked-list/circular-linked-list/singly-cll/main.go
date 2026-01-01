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
