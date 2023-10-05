package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "a", args: args{nums: []int{2, 5, 11, 15}, target: 13}, want: []int{0, 2}},
		{name: "a", args: args{nums: []int{2, 3, 11, 15}, target: 17}, want: []int{0, 3}},
		{name: "a", args: args{nums: []int{3, 11, 15}, target: 14}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSumHashTable(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "a", args: args{nums: []int{2, 5, 11, 15}, target: 13}, want: []int{0, 2}},
		{name: "a", args: args{nums: []int{2, 3, 11, 15}, target: 17}, want: []int{0, 3}},
		{name: "a", args: args{nums: []int{3, 11, 15}, target: 14}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumHashTable(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumHashTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
