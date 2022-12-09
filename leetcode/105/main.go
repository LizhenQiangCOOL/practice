package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	preStart, inStart := 0, 0
	preEnd, inEnd := len(preorder)-1, len(inorder)-1
	return build(&preorder, preStart, preEnd, &inorder, inStart, inEnd)
}

func build(preorder *[]int, preStart, preEnd int, inorder *[]int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootVal := (*preorder)[preStart]
	index := 0
	for i := inStart; i <= inEnd; i++ {
		if (*inorder)[i] == rootVal {
			index = i
			break
		}
	}
	leftSize := index - inStart
	root := &TreeNode{Val: rootVal}
	//递归构造左右子树
	root.Left = build(preorder, preStart+1, preStart+leftSize, inorder, inStart, index-1)
	root.Right = build(preorder, preStart+leftSize+1, preEnd, inorder, index+1, inEnd)
	return root
}
