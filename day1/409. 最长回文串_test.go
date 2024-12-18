package day1

import "testing"

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "bb",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
