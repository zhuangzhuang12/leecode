package day1

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	a := head
	b := head
	for i := 0; i < k; i++ {
		b = b.Next
		if b == nil {
			return head
		}
	}
	newHead := reverseList(a, b)
	a.Next = ReverseKGroup(b, k)
	return newHead
}

func reverseList(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur := a
	nxt := a
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
