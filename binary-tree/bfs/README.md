# Breadth First Search | BFS

BFS is the only tree algorithm that doesn't use recursion! It is the search through levels, you look all Nodes in each level and checks if they're what you're looking for... Normally we use BFS with another data structure called `Dequeue`, we use it because it's easir to implement in that way.

## Go Implementation:

```go
package main

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) Insert(data int) {
	if t.root == nil {
		t.root = &Node{
			value: data,
		}
	} else {
		insertRecursive(t.root, data)
	}
}

func (t *BinaryTree) BFS(data int) *Node {
	if t.root == nil {
		return nil
	}

    // I don't want to focus on the dequeue implementation, so I'm using an slice of pointers of Nodes
    // to make everything more simple...
    dequeue := make([]*Node, 0, 16)
	dequeue = append(dequeue, t.root)

	node := &Node{}
	for len(dequeue) > 0 {
		node = dequeue[0]
		dequeue = append(dequeue[:0], dequeue[1:]...) // This is just a way to delete the X element from my queue...

		if node.value == data {
			return node
		}

		if node.left != nil {
			dequeue = append(dequeue, node.left)
		}

		if node.right != nil {
			dequeue = append(dequeue, node.right)
		}
	}

	return nil
}

func insertRecursive(node *Node, data int) {
	if data > node.value {
		if node.right == nil {
			node.right = &Node{value: data}
			return
		}
		insertRecursive(node.right, data)
	} else {
		if node.left == nil {
			node.left = &Node{value: data}
			return
		}
		insertRecursive(node.left, data)
	}
}

func main() {
	tree := BinaryTree{}
	tree.Insert(2) // root
	tree.Insert(1) // root -> left
	tree.Insert(5) // root -> right
	tree.Insert(4) // root -> right -> left
	tree.Insert(6) // root -> right -> right

	if node := tree.BFS(6); node != nil {
		println(node.value)
	}
}
```
