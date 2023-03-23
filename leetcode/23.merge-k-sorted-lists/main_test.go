package main

import (
	"github.com/stretchr/testify/assert"
	"leetcode/util"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*util.ListNode
	}
	tests := []struct {
		name string
		args args
		want *util.ListNode
	}{
		// TODO: Add test cases.
		{
			name: "正常用例",
			args: args{
				lists: []*util.ListNode{
					util.ArrayToList([]int{1, 4, 5}),
					util.ArrayToList([]int{1, 3, 4}),
					util.ArrayToList([]int{2, 6}),
				},
			},
			want: util.ArrayToList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			name: "特殊用例",
			args: args{
				lists: []*util.ListNode{},
			},
			want: util.ArrayToList([]int{}),
		},
		{
			name: "特殊用例",
			args: args{
				lists: []*util.ListNode{
					util.ArrayToList([]int{}),
				},
			},
			want: util.ArrayToList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeKLists(tt.args.lists)
			wantL := util.ListToArray(tt.want)
			gotL := util.ListToArray(got)
			assert.EqualValuesf(t, wantL, gotL, "want %+v got %+v", wantL, gotL)
		})
	}
}
