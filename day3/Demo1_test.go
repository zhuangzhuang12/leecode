package day3

import "testing"

func Test_getFirst10Chars(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "一二三四五六七八九十",
				n: 8,
			},
			want: "1234567890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFirstNChars(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("getFirst10Chars() = %v, want %v", got, tt.want)
			}
		})
	}
}
