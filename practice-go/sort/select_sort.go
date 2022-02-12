package sort

// 选择排序是把最小的元素依次放到已排序列表中。
func SelectSort(A []int) {
	for i := 0; i < len(A); i++ {
		minIndex := i
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[minIndex] {
				minIndex = j
			}
		}
		A[i], A[minIndex] = A[minIndex], A[i]
	}
}
