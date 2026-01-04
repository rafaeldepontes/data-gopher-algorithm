# Depth First Search | DFS

DFS stands for Depth First Search, an algorithm to traverse a tree where you go through all the Nodes as deep as possible before bracktracking. The preorder, inorder and postorder are DFS variations! As we are know each algorithm use recursion, which implicitly uses a call stack and the DFS is not different!

But as said previously the preorder, inorder and postorder are variations! The real DFS doesn't check if the node is higher or lower... none of those things actually matter, because previously our tree was half-balanced! We can not say that our tree is fully balanced because we are not checking the depth...

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

func (t *BinaryTree) DFS(data int) *Node {
	return recursiveDFS(t.root, data)
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

func recursiveDFS(node *Node, data int) *Node {
	if node == nil {
		return nil
	}

	if node.value == data {
		return node
	}

	if n := recursiveDFS(node.left, data); n != nil {
		return n
	}

	if n := recursiveDFS(node.right, data); n != nil {
		return n
	}

	return nil
}

func main() {
	tree := BinaryTree{}
	tree.Insert(2) // root
	tree.Insert(1) // root -> left
	tree.Insert(5) // root -> right
	tree.Insert(4) // root -> right -> left
	tree.Insert(6) // root -> right -> right

	if node := tree.DFS(9); node != nil {
		println(node.value)
	}
}
```

