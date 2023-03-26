package main

import "leetcode/util"

var result []int

// https://leetcode.cn/problems/binary-tree-postorder-traversal/
func postorderTraversal(root *util.TreeNode) []int {
	result = make([]int, 0)
	traverse(root)
	return result
}
func traverse(root *util.TreeNode) {
	if root == nil {
		return
	}
	//前序位置
	traverse(root.Left)
	traverse(root.Right)
	//后序位置
	result = append(result, root.Val)
}
