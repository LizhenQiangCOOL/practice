package sort

func InsertionSort(A []int) {
	if len(A) < 2 {
		return
	}
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for ; i >= 0 && A[i] > key; i-- {
			A[i+i] = A[i]
		}
		A[i+1] = key
	}
}
