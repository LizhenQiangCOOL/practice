package main

import "leetcode/util"

var prev, first, second *util.TreeNode

func recoverTree(root *util.TreeNode) {
	prev, first, second = nil, nil, nil
	traverse(root)
	first.Val, second.Val = second.Val, first.Val
}

func traverse(node *util.TreeNode) {
	if node == nil {
		return
	}
	traverse(node.Left)
	if prev != nil && prev.Val > node.Val {
		if first == nil {
			first = prev
		}
		second = node
	}
	prev = node
	traverse(node.Right)
}
