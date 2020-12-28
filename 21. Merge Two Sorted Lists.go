package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	switch {
	case l1 == nil:
		return l2
	case l2 == nil:
		return l1
	case l1.Val < l2.Val:
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	default:
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
