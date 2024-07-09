package sort

// Parent 存储完全二叉树 e.g. parent index 0, left 1, right 2
// parent index 1, left 3, right, 4
func Parent(i int) int {
	return (i - 1) / 2
}
func Left(i int) int {
	return 2*(i+1) - 1
}
func Right(i int) int {
	return 2 * (i + 1)
}

func MaxHeapify(A []int, heapSize int, i int) {
	l := Left(i)
	r := Right(i)
	large := i
	if l < heapSize && A[l] > A[i] {
		large = l
	}
	if r < heapSize && A[r] > A[large] {
		large = r
	}
	if large != i {
		A[i], A[large] = A[large], A[i]
		MaxHeapify(A, heapSize, large)
	}
}

func BuildMaxHeap(A []int) {
	heapSize := len(A)
	for i := heapSize / 2; i >= 0; i-- {
		MaxHeapify(A, heapSize, i)
	}
}

func HeapSort(A []int) {
	BuildMaxHeap(A)
	heapSize := len(A)
	for i := len(A) - 1; i > 0; i-- {
		A[0], A[i] = A[i], A[0]
		heapSize -= 1
		MaxHeapify(A, heapSize, 0)
	}
}
