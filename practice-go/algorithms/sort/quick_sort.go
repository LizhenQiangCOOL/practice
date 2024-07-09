package sort

// QuickSort 快速排序
func QuickSort(A []int, left int, right int) {
	if left < right {
		index := Partition(A, left, right)
		QuickSort(A, left, index-1)
		QuickSort(A, index+1, right)
	}
}

// Partition A[left...right]，A[right] exist
func Partition(A []int, left int, right int) int {
	x := A[right]
	i := left - 1
	for j := left; j < right; j++ {
		if A[j] < x {
			i = i + 1
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[right] = A[right], A[i+1]
	// x value index
	return i + 1
}
