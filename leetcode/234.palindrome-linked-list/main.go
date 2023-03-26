package main

import "leetcode/util"

func isPalindrome(head *util.ListNode) bool {
	// base case
	if head == nil {
		return true
	}
	if head.Next == nil {
		return true
	}
	//快慢指针找到中间
	headNode := &util.ListNode{}
	headNode.Next = head
	fastP, slowP := headNode, headNode
	for fastP.Next != nil && fastP.Next.Next != nil {
		slowP = slowP.Next
		fastP = fastP.Next
		fastP = fastP.Next
	}
	another := reverseList(slowP.Next)
	p, anotherP := headNode.Next, another
	flag := true
	for p != nil && anotherP != nil {
		if p.Val != anotherP.Val {
			flag = false
			break
		}
		p = p.Next
		anotherP = anotherP.Next
	}
	slowP.Next = reverseList(another)
	return flag
}

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
