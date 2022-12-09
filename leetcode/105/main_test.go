package main

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "测试用例一",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: &TreeNode{3,
				&TreeNode{9, nil, nil},
				&TreeNode{20,
					&TreeNode{15, nil, nil},
					&TreeNode{7, nil, nil}}},
		},
		{
			name: "测试用例二",
			args: args{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: &TreeNode{-1, nil, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
