package day1

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	list := []*ListNode{}
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	i, j := 0, len(list)-1
	for i != j {
		list[i].Next = list[j]
		if i == j {
			break
		}
		i++
		list[j].Next = list[i]
		j--
	}
	list[i].Next = nil
}
