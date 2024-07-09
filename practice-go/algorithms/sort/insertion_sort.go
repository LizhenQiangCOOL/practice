package sort

// InsertionSort 插入排序
func InsertionSort(A []int) {
	for j := 0; j < len(A); j++ {
		key := A[j]
		i := j - 1
		// 往前找符合的位置，并腾位置
		for ; i >= 0 && A[i] > key; i-- {
			A[i+1] = A[i]
		}
		// 找到符合的位置
		A[i+1] = key
	}
}
