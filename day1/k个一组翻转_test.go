package day1

import (
	"reflect"
	"testing"
)

func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, val := range vals {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return dummy.Next
}
func TestReverseKGroup(t *testing.T) {
	head := buildList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: head,
				k:    3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
