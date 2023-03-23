package main

import "leetcode/util"

func partition(head *util.ListNode, x int) *util.ListNode {
	resultHead := &util.ListNode{}
	helpHead := &util.ListNode{}
	resultP, helpP, p := resultHead, helpHead, head
	for p != nil {
		if p.Val < x {
			resultP.Next = p
			resultP = resultP.Next
			p = p.Next
			resultP.Next = nil
		} else {
			helpP.Next = p
			helpP = helpP.Next
			p = p.Next
			helpP.Next = nil
		}
	}
	resultP.Next = helpHead.Next
	return resultHead.Next
}
