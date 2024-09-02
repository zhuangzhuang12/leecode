package day1

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	c := reverse(a, b)
	a.Next = ReverseKGroup(b, k)
	return c
}
func reverse(a, b *ListNode) *ListNode {
	var pre, cur, nxt *ListNode
	pre = nil
	cur = a
	nxt = a
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
