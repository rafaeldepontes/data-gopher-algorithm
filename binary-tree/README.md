# Binary Tree

A tree is basically just another type of graph, but for the sake of learning... We will considere both as different things.

A Tree is a data structures that always will have a Root Node, the tree has a depth and the children's Node. A beatiful Tree to look at could be the one outside, but as we are talking about data structure the one that I'm talking about here is the directory tree... If you open any folder/directory in our desktop/notebook you can see you files in some type of order:

```bash
c:/
    photos/
    videos/
    package.json
    .git/
    README.md
```

This my friends are a tree, not a binary tree... just a Tree.

A Binary Tree is the same thing but with limitations... Like for example, as the name suggest, a Binary Tree is a tree that each Node has only up to two children. So If we look at the example above... `C:` it's not a binary tree since it has five contents.

## Golang Implementation:

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
		t.insertRecursive(t.root, data)
	}
}

func (t *BinaryTree) Search(data int) *Node {
	return t.searchRecursive(t.root, data)
}

func (t *BinaryTree) searchRecursive(node *Node, data int) *Node {
	if node == nil {
		return nil
	}

	if node.value == data {
		return node
	}

	if node.value < data {
		if n := t.searchRecursive(node.right, data); n != nil {
			return n
		}
	} else {
		if n := t.searchRecursive(node.left, data); n != nil {
			return n
		}
	}

	return nil
}

func (t *BinaryTree) insertRecursive(node *Node, data int) {
	if data > node.value {
		if node.right == nil {
			node.right = &Node{value: data}
			return
		}
		t.insertRecursive(node.right, data)
	} else {
		if node.left == nil {
			node.left = &Node{value: data}
			return
		}
		t.insertRecursive(node.left, data)
	}
}

func main() {
	tree := BinaryTree{}
	tree.Insert(2) // root
	tree.Insert(1) // root -> left
	tree.Insert(5) // root -> right
	tree.Insert(4) // root -> right -> left
	tree.Insert(6) // root -> right -> right

	// println(tree.root.value, tree.root.left.value, tree.root.right.value, tree.root.right.left.value, tree.root.right.right.value)

	if node := tree.Search(6); node != nil {
		println(node.value)
	}

	if node := tree.Search(9); node != nil {
		println(node.value)
	}
}
```

## Preorder Traversal

Well now that we have our tree implemented, we can just use a little bit of logic here to print our tree... don't you think?

The Preorder Traversal is a depth-first method, it visits each node in the order `Root`, `Left` and finally `Right`. So if we have a tree like this:

```bash
        2
    /       \
   1        4
          /   \
        5       6
```

The Preorder result would be:

```log
*2* 1 5 4 6
```

> A easy way to remember is preorder the root will always be the first thing...

## Go Implementation:

```go
func (t *BinaryTree) PrintTreePreorder() {
	print("Tree: ")
	t.preorderTraversal(t.root)
	println()
}

func (t *BinaryTree) preorderTraversal(node *Node) {
	if node == nil {
		return
	}

	print(node.value, " ")
	if node.left != nil {
		t.preorderTraversal(node.left)
	}

	if node.right != nil {
		t.preorderTraversal(node.right)
	}
}

func main() {
    tree := BinaryTree{}
	tree.Insert(2) // root
	tree.Insert(1) // root -> left
	tree.Insert(5) // root -> right
	tree.Insert(4) // root -> right -> left
	tree.Insert(6) // root -> right -> right
    // ...
    tree.PrintTreePreorder()
}
```

## Postorder Traversal

Knowing the Preorder, we can easilly build our own postorder! Like the preorder where the root is always at the beginning, the postorder the root is always at the end.

So for the same tree, the postorder would be:

```log
1 4 6 5 *2*
```

## Go Implementation:

```go
func (t *BinaryTree) PrintTreePostorder() {
	print("Tree: ")
	t.postorderTraversal(t.root)
	println()
}

func (t *BinaryTree) postorderTraversal(node *Node) {
	if node == nil {
		return
	}

	if node.left != nil {
		t.postorderTraversal(node.left)
	}

	if node.right != nil {
		t.postorderTraversal(node.right)
	}

	print(node.value, " ")
}

func main() {
    tree := BinaryTree{}
	tree.Insert(2) // root
	tree.Insert(1) // root -> left
	tree.Insert(5) // root -> right
	tree.Insert(4) // root -> right -> left
	tree.Insert(6) // root -> right -> right
    // ...
    tree.PrintTreePostorder()
}
```

## Inorder Traversal

And finally for the Inorder traversal, the easiest one, the root somewhere in the middle

```log
1 *2* 5 4 6
```

## Go Implementation:

```go
func (t *BinaryTree) PrintTreeInorder() {
	print("Tree: ")
	t.inorderTraversal(t.root)
	println()
}

func (t *BinaryTree) inorderTraversal(node *Node) {
	if node == nil {
		return
	}

	if node.left != nil {
		t.preorderTraversal(node.left)
	}

	print(node.value, " ")

	if node.right != nil {
		t.preorderTraversal(node.right)
	}
}

func main() {
    tree := BinaryTree{}
	tree.Insert(2) // root
	tree.Insert(1) // root -> left
	tree.Insert(5) // root -> right
	tree.Insert(4) // root -> right -> left
	tree.Insert(6) // root -> right -> right
    // ...
    tree.PrintTreeInorder()
}
```
