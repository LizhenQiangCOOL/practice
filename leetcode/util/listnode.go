package util

type ListNode struct {
	Val  int
	Next *ListNode
}

// 数组转换成单向链表
func ArrayToList(s []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, v := range s {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return head.Next
}

// 单向链表转换成数组
func ListToArray(l *ListNode) []int {
	s := make([]int, 0)
	p := l
	for p != nil {
		s = append(s, p.Val)
		p = p.Next
	}
	return s
}

// 实现最大堆，即实现 container 下 heap
type MaxHeapList []*ListNode

func (h MaxHeapList) Len() int {
	return len(h)
}
func (h MaxHeapList) Less(i, j int) bool {
	// 返回ture则不改变i，j元素位置
	return h[i].Val > h[j].Val
}
func (h MaxHeapList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeapList) Push(v *ListNode) {
	*h = append(*h, v)
}
func (h *MaxHeapList) Pop() (v *ListNode) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

// 实现最小堆，即实现 container 下 heap
type MinHeapList []*ListNode

func (h MinHeapList) Len() int {
	return len(h)
}
func (h MinHeapList) Less(i, j int) bool {
	// 返回ture则不改变i，j元素位置
	return h[i].Val < h[j].Val
}
func (h MinHeapList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeapList) Push(v any) {
	*h = append(*h, v.(*ListNode))
}
func (h *MinHeapList) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}
