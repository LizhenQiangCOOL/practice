package main

import (
	"container/heap"
	"leetcode/util"
)

func mergeKLists(lists []*util.ListNode) *util.ListNode {
	head := &util.ListNode{}
	minHeap := &util.MinHeapList{}
	heap.Init(minHeap)
	headP := head
	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}
	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(*util.ListNode)
		headP.Next = node
		if node.Next != nil {
			//说明该链表还存在元素需要继续放入最小堆
			heap.Push(minHeap, node.Next)
		}
		headP = headP.Next
	}
	return head.Next
}
