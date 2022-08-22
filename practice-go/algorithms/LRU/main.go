package main

import "fmt"

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (list *DoublyLinkedList) AppendToFront(node *Node) {
	node.next = list.head.next
	if node.next != nil {
		node.next.pre = node
	} else {
		// 链表插入前为空
		list.tail.next = node
	}
	list.head.next = node
	node.pre = list.head
}

func (list *DoublyLinkedList) MoveToFront(node *Node) {
	// 如果node为最后一个节点
	if node.next == nil {
		// 剩余1个节点或者为空
		if list.head.next == list.tail.next {
			return
		} else {
			// 需要处理尾指针
			list.tail.next = node.pre
		}
	} else {

	}
	node.pre.next = node.next
}

func main() {
	fmt.Println("1111111111")
}
