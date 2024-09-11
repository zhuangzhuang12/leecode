package day1

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	node := find(dummy, n)
	node.Next = node.Next.Next
	return dummy.Next
}

func find(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
