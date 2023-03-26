package main

import (
	"fmt"
	"leetcode/util"
)

func trimBST(root *util.TreeNode, low int, high int) *util.TreeNode {
	// 构建空父亲节点，避免根节点要删除情况
	rootParents := &util.TreeNode{Val: low}
	rootParents.Left = root

	traverse(root, rootParents, true, low, high)
	return rootParents.Left
}

// flag 为true 表示 root为parents的左子树，否则root为paranets的右子树
func traverse(root *util.TreeNode, parents *util.TreeNode, flag bool, low int, high int) {
	if root == nil {
		return
	}
	//前序位置，减少要处理的节点
	if root.Val < low {
		//说明当前节点只要右边
		if flag {
			parents.Left = root.Right
		} else {
			parents.Right = root.Right
		}
		traverse(root.Right, parents, flag, low, high)
		return
	}
	if root.Val > high {
		if flag {
			parents.Left = root.Left
		} else {
			parents.Right = root.Left
		}
		traverse(root.Left, parents, flag, low, high)
		return
	}
	traverse(root.Left, root, true, low, high)
	traverse(root.Right, root, false, low, high)
	//后序位置
}

func traverseNew(root *util.TreeNode, low, high int) *util.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return traverseNew(root.Right, low, high)
	}
	if root.Val > high {
		return traverseNew(root.Left, low, high)
	}
	root.Left = traverseNew(root.Left, low, high)
	root.Right = traverseNew(root.Right, low, high)
	return root
}

func main() {
	a := &util.TreeNode{Val: 3}
	b := &util.TreeNode{Val: 1}
	c := &util.TreeNode{Val: 4}
	d := &util.TreeNode{Val: 2}
	a.Left = b
	a.Right = c
	b.Left = d
	low, high := 3, 4
	result := trimBST(a, low, high)
	fmt.Println(result)
}
