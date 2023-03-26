package main

import "leetcode/util"

var result [][]int
var trackDeepth int

// https://leetcode.cn/problems/binary-tree-level-order-traversal/
// 递归解法
func levelOrder(root *util.TreeNode) [][]int {
	result = make([][]int, 0)
	trackDeepth = -1
	traverse(root, 0)
	return result
}

func traverse(root *util.TreeNode, deepth int) {
	if root == nil {
		return
	}
	//说明当前是新的一个数组
	if deepth > trackDeepth {
		trackDeepth += 1
		track := make([]int, 1)
		track[0] = root.Val
		result = append(result, track)
	} else {
		//说明非新的数组
		if len(result) > deepth {
			track := result[deepth]
			track = append(track, root.Val)
			result[deepth] = track
		}
	}

	traverse(root.Left, deepth+1)
	traverse(root.Right, deepth+1)
}

// 非递归解法
func levelOrderY(root *util.TreeNode) [][]int {
	resultY := make([][]int, 0)
	queue := make([]*util.TreeNode, 0)
	queueDeepth := make([]int, 0)
	queue = append(queue, root)
	queueDeepth = append(queueDeepth, 0)
	trackDeepthY := -1
	for len(queue) > 0 && len(queueDeepth) > 0 {
		node := queue[0]
		nodeDeepth := queueDeepth[0]
		queue = queue[1:]
		queueDeepth = queueDeepth[1:]
		if nodeDeepth > trackDeepthY {
			trackDeepthY += 1
			track := make([]int, 0)
			track = append(track, node.Val)
			resultY = append(resultY, track)
		} else {
			if len(resultY) > nodeDeepth {
				track := resultY[nodeDeepth]
				track = append(track, node.Val)
				resultY[nodeDeepth] = track
			}
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
			queueDeepth = append(queueDeepth, nodeDeepth+1)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			queueDeepth = append(queueDeepth, nodeDeepth+1)
		}
	}
	return resultY
}
