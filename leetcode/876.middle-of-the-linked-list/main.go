package main

import "leetcode/util"

func middleNode(head *util.ListNode) *util.ListNode {
	headFastP, headSlowP := head, head
	for !(headFastP == nil || headFastP.Next == nil) {
		// 快指针肯定能走第一步
		headFastP = headFastP.Next
		headFastP = headFastP.Next
		headSlowP = headSlowP.Next
	}
	return headSlowP
}
