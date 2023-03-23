package main

import "leetcode/util"

func getIntersectionNode(headA, headB *util.ListNode) *util.ListNode {
	headANode, headBNode := &util.ListNode{Next: headA}, &util.ListNode{Next: headB}
	headAP, headBP := headANode, headBNode
	flagA, flagB := 0, 0
	for headAP != headBP {
		if flagA > 1 || flagB > 1 {
			//不存在环
			return nil
		}
		if headAP.Next == nil {
			//走完一次需要走B
			headAP = headBNode
			flagA += 1
		} else {
			headAP = headAP.Next
		}
		if headBP.Next == nil {
			headBP = headANode
			flagB += 1
		} else {
			headBP = headBP.Next
		}
	}
	return headAP
}
