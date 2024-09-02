package day1

import "testing"

func TestSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("SubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
