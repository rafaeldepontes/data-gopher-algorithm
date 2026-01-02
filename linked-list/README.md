# Linked List

A linear data structure that contains a node if X information and a next, tail or head that stores the reference to the subsequent node.

The entry point of a linked list is called HEAD (normally a pointer to the first value of the linked list), usually we will or can see things like "NEXT" which in this case would be the reference to the next node on our list... If it does not exist then its NULL (or nil in this case.), we can also have the PREVIOUS (or PREV) that is a reference to the previous node on our list and there are cases which we can encounter a TAIL, the TAIL normally refers to the reference to the LAST node in or list.

There are **three** types of Linked Lists:

- Singly Linked List
- Doubly Linked List
- Circular Linked List

Besides the types, there are some concepts that I find interesting to know... One of them is how to reverse a linked list!

To reverse a linked list depends on what type of linked list we are handling, if it's a singly linked list: You can create a new linked list, maybe calling `reversedLinkedList`, do a loop while the list is not empty, store the value of the next node into a `temp` variable, then you say that your next node receives your `reversedLinkedList`, then your `reversedLinkedList` receives the current `head/list` and FINALLY the `head/list` receives the `temp` node!

Cool talking and all, but let's see in practice:

## Go Implementation:

This LeetCode problem can be easily solved by our strategy:

### 206. Reverse Linked List

Given the `head` of a singly linked list, reverse the list, and return _the reversed_ list.

Example 1:

- Input: head = [1,2,3,4,5]
- Output: [5,4,3,2,1]

### Code:

```go
package main

 type ListNode struct {
     Val int
     Next *ListNode
 }

func reverseList(head *ListNode) *ListNode {
    var ans *ListNode

    for head != nil {
		nextNode := head.Next
		head.Next = ans
		ans = head
		head = nextNode
    }

    return head
}
```
