package main

import (
	"leetcode/util"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *util.TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "测试用例一",
			args: args{
				root: &util.TreeNode{3,
					&util.TreeNode{1, nil, &util.TreeNode{2, nil, nil}},
					&util.TreeNode{4, nil, nil}},
			},
			want: int(1),
		},
		{
			name: "测试用例二	",
			args: args{
				root: &util.TreeNode{5,
					&util.TreeNode{3,
						&util.TreeNode{2,
							&util.TreeNode{1, nil, nil}, nil},
						&util.TreeNode{4, nil, nil}},
					&util.TreeNode{6, nil, nil}},
			},
			want: int(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
