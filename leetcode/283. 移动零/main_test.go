package main

import (
	"fmt"
	"reflect"
	"testing"
)

func MoveZeroes(nums []int) {
	// nums 为 切片, 为引用类型
	var n int = len(nums)
	var fast, slow int = 0, 0
	for fast < n {
		if nums[fast] != 0 && nums[slow] == 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			fast++
			slow++
			continue
		}
		if fast == slow {
			fast++
			continue
		}
		if nums[fast] == 0 {
			fast++
		}
		if nums[slow] != 0 {
			slow++
		}
	}
}

func MoveZeroes_leetcode(nums []int) {
	var n int = len(nums)
	var fast, slow int = 0, 0
	for fast < n {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}

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
