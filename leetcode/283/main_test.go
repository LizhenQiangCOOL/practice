package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		//正常例子
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		//特殊例子
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		//边界例子
		{[]int{0, 1, 2, 3, 0}, []int{1, 2, 3, 0, 0}},
		{[]int{1, 0, 0, 0, 2}, []int{1, 2, 0, 0, 0}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.nums)
		t.Run(testname, func(t *testing.T) {
			MoveZeroes(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("got %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
