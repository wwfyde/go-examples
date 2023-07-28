package helloworld

import "testing"

func Test_hello(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "世界", args: args{"世界"}, want: "hello, 世界!"},
		{name: "empty", args: args{""}, want: "hello, !"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hello(tt.args.in); got != tt.want {
				t.Errorf("hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
