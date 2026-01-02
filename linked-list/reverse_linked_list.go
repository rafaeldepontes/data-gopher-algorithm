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