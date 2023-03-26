package main

import "leetcode/util"

var res = 0
var depth = 0

func maxDepth(root *util.TreeNode) int {
	//有变量依赖
	res, depth = 0, 0
	traverse(root)
	return res
}

// 遍历解决问题，回溯
func traverse(root *util.TreeNode) {
	if root == nil {
		return
	}
	//前序位置，遍历节点前
	depth++

	if root.Left == nil && root.Right == nil {
		//到达叶子节点，更新最大深度
		if depth > res {
			res = depth
		}
	}
	traverse(root.Left)
	traverse(root.Right)
	//后序位置，便利完节点回来
	depth--
}

// 分解问题解决，分解成子问题
func maxDepth2(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	//计算左子树和右子树
	leftMax := maxDepth2(root.Left)
	rightMax := maxDepth2(root.Right)
	res := max(leftMax, rightMax) + 1
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
