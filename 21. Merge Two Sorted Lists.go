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

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := new(ListNode)
	prev := prehead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}
	return prehead.Next
}
