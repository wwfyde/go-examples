package hello

import "testing"

func Test_hello(t *testing.T) {
	// test Cases
	testCases := []struct {
		name, in, want string
	}{
		{"empty", "", "Hello, !"},
		{"世界", "世界", "Hello, 世界!"},
		{"world", "world", "Hello, world!"},
		{"mixed", "某 and some", "Hello, 某 and some!"},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.in); got != tt.want {
				t.Errorf("Hello = %s, want: %s", got, tt.want)
			}
		})
	}

}
