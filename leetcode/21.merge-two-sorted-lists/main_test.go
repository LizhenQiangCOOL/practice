package main

import (
	"github.com/stretchr/testify/assert"
	"leetcode/util"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *util.ListNode
		list2 *util.ListNode
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
				list1: util.ArrayToList([]int{1, 2, 4}),
				list2: util.ArrayToList([]int{1, 3, 4}),
			},
			want: util.ArrayToList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "特殊用例-全空",
			args: args{
				list1: util.ArrayToList([]int{}),
				list2: util.ArrayToList([]int{}),
			},
			want: util.ArrayToList([]int{}),
		}, {
			name: "特殊用例-一边为空",
			args: args{
				list1: util.ArrayToList([]int{1, 2, 4}),
				list2: util.ArrayToList([]int{}),
			},
			want: util.ArrayToList([]int{1, 2, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.list1, tt.args.list2)
			wantL := util.ListToArray(tt.want)
			gotL := util.ListToArray(got)
			assert.EqualValuesf(t, wantL, gotL, "want %+v got %+v", wantL, gotL)
		})
	}
}
