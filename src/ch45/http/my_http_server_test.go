package my_http

import (
	"fmt"
	"testing"
)

func Fib(n int) string {
	res := ""
	for a, b := 0, 1; b < n; a, b = b, a+b {
		res += fmt.Sprintf("%d, ", b)
	}
	return res
}

func TestFib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				n: 10,
			},
			want: "1, 1, 2, 3, 5, 8, ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
