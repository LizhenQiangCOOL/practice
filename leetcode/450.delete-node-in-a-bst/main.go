package main

import "leetcode/util"

func deleteNode(root *util.TreeNode, key int) *util.TreeNode {
	paranets := &util.TreeNode{}
	paranets.Left = root
	traverse(root, paranets, true, key)
	return paranets.Left
}

// flag 为 true 表示左子树
func traverse(root *util.TreeNode, parents *util.TreeNode, flag bool, key int) {
	if root == nil {
		return
	}
	if root.Val < key {
		//说明左边都不需要遍历了
		traverse(root.Right, root, false, key)
		return
	}
	if root.Val > key {
		//说明右边都不需要遍历了
		traverse(root.Left, root, true, key)
		return
	}
	if root.Val == key {
		//要进行操作
		leftNode := getLeftNode(root.Right)
		if leftNode != nil {
			//存在该节点
			leftNode.Left = root.Left
			if flag {
				parents.Left = root.Right
			} else {
				parents.Right = root.Right
			}
		} else {
			//不存在右节点
			if flag {
				parents.Left = root.Left
			} else {
				parents.Right = root.Left
			}
		}
		// 由于节点值是唯一的所以可以return
		return
	}
	traverse(root, parents, true, key)
	traverse(root, parents, false, key)
}

// 找到一个树的最左边
func getLeftNode(root *util.TreeNode) *util.TreeNode {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}
