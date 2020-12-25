package main

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var curr = head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}
