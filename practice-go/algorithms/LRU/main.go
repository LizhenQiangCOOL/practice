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
		// 处理后一个节点的前驱
		node.next.pre = node.pre
	}
	// 处理前一个节点的后继
	node.pre.next = node.next
	// 移动到前面，也要处理node后继和前驱
	node.next = list.head.next
	list.head.next.pre = node

	list.head.next = node
	node.pre = list.head
}

func (list *DoublyLinkedList) RemoveFromTail() {
	if list.head.next == list.tail.next {
		list.head.next = nil
		list.tail.next = nil
	} else {
		lastNode := list.tail.next
		lastNode.pre.next = nil
		list.tail.next = list.tail.next.pre
		lastNode.pre = nil
	}
}

type LRUCache struct {
	maxSize    int
	size       int
	lookUp     map[int]*Node
	linkedList *DoublyLinkedList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		maxSize: capacity,
		size:    0,
		lookUp:  make(map[int]*Node),
		linkedList: &DoublyLinkedList{
			head: &Node{},
			tail: &Node{},
		},
	}
}

func (cache *LRUCache) get(key int) int {
	node, pre := cache.lookUp[key]
	if !pre {
		return -1
	}
	cache.linkedList.MoveToFront(node)
	return node.value
}

func (cache *LRUCache) set(key, value int) {
	node, pre := cache.lookUp[key]
	if !pre {
		if cache.size == cache.maxSize {
			lastNode := cache.linkedList.tail.next
			delete(cache.lookUp, lastNode.key)
			cache.linkedList.RemoveFromTail()
		} else {
			cache.size += 1
		}
		newNode := &Node{
			key:   key,
			value: value,
		}
		cache.linkedList.AppendToFront(newNode)
		cache.lookUp[key] = newNode
	} else {
		node.value = value
		cache.linkedList.MoveToFront(node)
	}
}

func main() {
	fmt.Println("1111111111")
}
