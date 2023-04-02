package main

import (
	"leetcode/util"
)

var sum int

func convertBST(root *util.TreeNode) *util.TreeNode {
	sum = 0
	traverse(root)
	return root
}

func traverse(root *util.TreeNode) {
	if root == nil {
		return
	}

	traverse(root.Right)
	//先搜索右边,由于唯一值+二叉搜索树的性质 => 右边都比中间的大
	//得出当前结果
	sum += root.Val
	root.Val = sum

	traverse(root.Left)
}
