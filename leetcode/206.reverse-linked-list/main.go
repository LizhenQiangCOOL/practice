package main

import "leetcode/util"

func reverseList(head *util.ListNode) *util.ListNode {
	headNew := &util.ListNode{}
	headResult := &util.ListNode{}
	headNew.Next = head
	for headNew.Next != nil {
		tmp := headNew.Next
		headNew.Next = headNew.Next.Next

		// 插到新链表
		tmp.Next = headResult.Next
		headResult.Next = tmp
	}
	return headResult.Next
}
