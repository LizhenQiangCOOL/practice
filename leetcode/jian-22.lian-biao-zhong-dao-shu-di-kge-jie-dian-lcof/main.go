package main

import "leetcode/util"

func getKthFromEnd(head *util.ListNode, k int) *util.ListNode {
	// 检查链表是否是有那么长
	headFastP, headSlowP := head, head
	for i := 0; i < k; i++ {
		if headFastP != nil {
			headFastP = headFastP.Next
		} else {
			return nil
		}
	}
	for headFastP != nil {
		headFastP = headFastP.Next
		headSlowP = headSlowP.Next
	}
	return headSlowP
}
