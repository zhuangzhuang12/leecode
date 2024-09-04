package day1

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	next, cur := head, head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func ReverseList01(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseList01(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
