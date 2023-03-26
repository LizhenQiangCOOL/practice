package main

import "leetcode/util"

var maxSum = 0

func diameterOfBinaryTree(root *util.TreeNode) int {
	maxSum = 0
	maxDepth(root)
	return maxSum
}

func maxDepth(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	// 后序位置，已经知道子树的数据
	myMax := leftMax + rightMax
	maxSum = max(myMax, maxSum)
	return 1 + max(leftMax, rightMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
