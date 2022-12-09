package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	res, rank := 0, 0
	var traverse func(root *TreeNode, k int)
	traverse = func(root *TreeNode, k int) {
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
