package main

import "leetcode/util"

var max = 0
var result = []int{}

func findMode(root *util.TreeNode) []int {
	max = 0
	result = make([]int, 0)
	traverse(root)
	return result
}

func traverse(root *util.TreeNode) {
	if root == nil {
		return
	}
	knum := numberTraverse(root, root.Val)
	if knum > max {
		max = knum
		result = make([]int, 0)
	}
	if knum == max {
		result = append(result, root.Val)
	}
	traverse(root.Left)
	traverse(root.Right)
}

func numberTraverse(root *util.TreeNode, k int) int {
	if root == nil {
		return 0
	}
	leftNum := numberTraverse(root.Left, k)
	rightNum := numberTraverse(root.Right, k)
	addnum := 0
	if root.Val == k {
		addnum = 1
	}
	return leftNum + rightNum + addnum
}
