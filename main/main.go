package main

import (
	"leecode/day1"
)

func main() {
	a := &day1.ListNode{Val: 1, Next: &day1.ListNode{Val: 2, Next: &day1.ListNode{Val: 3, Next: &day1.ListNode{Val: 4, Next: &day1.ListNode{Val: 5, Next: nil}}}}}
	node := day1.ReverseList(a)
	for node != nil {
		println(node.Val)
		node = node.Next
	}
}

//func main() {
//	nums := []int{1, 2, 3}
//
//	fmt.Println(day1.Permute(nums))
//}
