package main

import "leetcode/util"

func mergeTwoLists(list1 *util.ListNode, list2 *util.ListNode) *util.ListNode {
	//生成空节点
	head := &util.ListNode{}
	newListP := head
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			newListP.Next = p1
			p1 = p1.Next
		} else {
			newListP.Next = p2
			p2 = p2.Next
		}
		newListP = newListP.Next
	}
	if p1 != nil {
		newListP.Next = p1
	}
	if p2 != nil {
		newListP.Next = p2
	}
	return head.Next
}
