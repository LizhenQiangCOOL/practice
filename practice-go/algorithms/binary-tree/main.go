package binary_tree

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func preorderTreeWalk(root *TreeNode) {
	if root != nil {
		preorderTreeWalk(root.left)

	}
}
