package main

import (
	"github.com/stretchr/testify/assert"
	"leetcode/util"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		head *util.ListNode
		k    int
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
				head: util.ArrayToList([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
			want: util.ArrayToList([]int{4, 5}),
		},
		{
			name: "特殊用例",
			args: args{
				head: util.ArrayToList([]int{1, 2, 3, 4, 5}),
				k:    6,
			},
			want: util.ArrayToList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getKthFromEnd(tt.args.head, tt.args.k)
			wantL, gotL := util.ListToArray(tt.want), util.ListToArray(got)
			assert.EqualValuesf(t, wantL, gotL, "want %+v got %+v", wantL, gotL)
		})
	}
}
