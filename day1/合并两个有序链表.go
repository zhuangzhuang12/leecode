package day1

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur := &ListNode{}
	dummy := &ListNode{}
	dummy.Next = cur
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next.Next
}
