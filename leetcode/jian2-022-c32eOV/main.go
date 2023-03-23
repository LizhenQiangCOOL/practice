package main

import (
	"fmt"
	"leetcode/util"
)

func detectCycle(head *util.ListNode) *util.ListNode {
	nHead := &util.ListNode{}
	nHead.Next = head
	fastP, slowP := nHead, nHead
	for fastP != slowP || fastP == nHead {
		if fastP.Next == nil {
			return nil
		}
		if fastP.Next.Next == nil {
			return nil
		}
		// 正常遍历
		fastP = fastP.Next
		fastP = fastP.Next
		slowP = slowP.Next
	}

	newP := nHead
	for newP != slowP {
		newP = newP.Next
		slowP = slowP.Next
	}
	return newP
}

func main() {
	a := &util.ListNode{Val: 3}
	b := &util.ListNode{Val: 2}
	c := &util.ListNode{Val: 0}
	d := &util.ListNode{Val: 4}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = b
	head := a
	result := detectCycle(head)
	fmt.Printf("result val: %v\n", result)
}
