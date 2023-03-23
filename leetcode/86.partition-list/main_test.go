package main

import (
	"github.com/stretchr/testify/assert"
	"leetcode/util"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *util.ListNode
		x    int
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
				head: util.ArrayToList([]int{1, 4, 3, 2, 5, 2}),
				x:    3,
			},
			want: util.ArrayToList([]int{1, 2, 2, 4, 3, 5}),
		},
		{
			name: "正常用例",
			args: args{
				head: util.ArrayToList([]int{2, 1}),
				x:    2,
			},
			want: util.ArrayToList([]int{1, 2}),
		},
		{
			name: "特殊用例-不命中",
			args: args{
				head: util.ArrayToList([]int{1, 2, 3, 4}),
				x:    5,
			},
			want: util.ArrayToList([]int{1, 2, 3, 4}),
		}, {
			name: "特殊用例-不命中",
			args: args{
				head: util.ArrayToList([]int{1, 2, 3, 4}),
				x:    0,
			},
			want: util.ArrayToList([]int{1, 2, 3, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.args.head, tt.args.x)
			wantL, gotL := util.ListToArray(tt.want), util.ListToArray(got)
			assert.EqualValuesf(t, wantL, gotL, "want %+v got %+v", wantL, gotL)
		})
	}
}
