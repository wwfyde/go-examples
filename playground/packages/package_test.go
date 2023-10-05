package packages

import (
	"fmt"
	"math"
	"testing"
)

func TestCircle_面积(t *testing.T) {
	type fields struct {
		Point  Point
		Radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{"0", fields{Radius: 0}, 0},
		{"1", fields{Radius: 1}, math.Pi * 1},
		{"2", fields{Radius: 2}, math.Pi * 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Circle{
				Point:  tt.fields.Point,
				Radius: tt.fields.Radius,
			}
			if got := c.面积(); got != tt.want {
				t.Errorf("面积() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine_Length(t *testing.T) {
	type fields struct {
		Start Point
		End   Point
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{"empty", fields{Start: Point{0, 0}, End: Point{0, 0}}, 0},
		{"1", fields{Start: Point{0, 0}, End: Point{1, 1}}, math.Sqrt(2)},
		{"1", fields{Start: Point{0, 0}, End: Point{2, 2}}, math.Sqrt(8)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Line{
				Start: tt.fields.Start,
				End:   tt.fields.End,
			}
			if got := l.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLength(t *testing.T) {
	type args struct {
		s Point
		e Point
	}
	tests := []struct {
		name       string
		args       args
		wantLength float64
	}{
		// TODO: Add test cases.
		{name: "empty", args: args{Point{0, 0}, Point{0, 0}}, wantLength: 0},
		{name: "1", args: args{Point{0, 0}, Point{1, 1}}, wantLength: math.Sqrt(2)},
		{name: "{1, 2}", args: args{Point{0, 0}, Point{1, 2}}, wantLength: math.Sqrt(5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLength := Length(tt.args.s, tt.args.e); gotLength != tt.wantLength {
				t.Errorf("Length() = %v, want %v", gotLength, tt.wantLength)
			}
		})
	}
}

func ExampleLength() {
	fmt.Println(Length(Point{}, Point{X: 1, Y: 2}))
	fmt.Println(Length(Point{}, Point{X: 2, Y: 5}))
}
