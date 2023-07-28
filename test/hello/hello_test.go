package hello

import "testing"

func Test_hello(t *testing.T) {
	// test Cases
	testCases := []struct {
		name, in, want string
	}{
		{"empty", "", "hello, !"},
		{"世界", "世界", "hello, 世界!"},
		{"world", "world", "hello, world!"},
		{"mixed", "某 and some", "hello, 某 and some!"},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := hello(tt.in); got != tt.want {
				t.Errorf("hello = %s, want: %s", got, tt.want)
			}
		})
	}

}
