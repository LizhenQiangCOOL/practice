package main

import (
	"github.com/stretchr/testify/assert"
	"leetcode/util"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *util.ListNode
	}
	tests := []struct {
		name string
		args args
		want *util.ListNode
	}{
		// TODO: Add test cases.
		{
			name: "正常例子",
			args: args{head: util.ArrayToList([]int{1, 2, 3})},
			want: util.ArrayToList([]int{2, 3}),
		},
		{
			name: "正常例子",
			args: args{head: util.ArrayToList([]int{1, 2})},
			want: util.ArrayToList([]int{2}),
		},
		{
			name: "正常例子",
			args: args{head: util.ArrayToList([]int{1})},
			want: util.ArrayToList([]int{1}),
		},
		{
			name: "正常例子",
			args: args{head: util.ArrayToList([]int{})},
			want: util.ArrayToList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := middleNode(tt.args.head)
			wantL, gotL := util.ListToArray(tt.want), util.ListToArray(got)
			assert.EqualValuesf(t, wantL, gotL, "want %+v got %+v", wantL, gotL)
		})
	}
}
