package main

import "leetcode/util"

func kthSmallest(root *util.TreeNode, k int) int {
	res, rank := 0, 0
	var traverse func(root *util.TreeNode, k int)
	traverse = func(root *util.TreeNode, k int) {
		if root == nil {
			return
		}
		traverse(root.Left, k)
		rank += 1
		if k == rank {
			res = root.Val
			return
		}
		traverse(root.Right, k)
	}
	traverse(root, k)
	return res
}
