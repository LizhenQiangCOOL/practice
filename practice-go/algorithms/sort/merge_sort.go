package sort

import "math"

func Merge(A []int, left int, mid int, right int) {
	// n1 第一个数组长度，n2 第二个数组长度
	n1 := mid - left + 1
	// 满足right-mid为第二个数组长度即可
	n2 := right - mid
	//let L[0..n1] and R[0...n2] be new arrays
	L := make([]int, n1+1)
	R := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		L[i] = A[left+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[mid+i+1]
	}
	// 增加哨兵
	L[n1] = math.MaxInt
	R[n2] = math.MaxInt
	i, j := 0, 0
	for k := left; k <= right; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i += 1
		} else {
			A[k] = R[j]
			j += 1
		}
	}
}

// MergeSort A[left..right]
func MergeSort(A []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		MergeSort(A, left, mid)
		MergeSort(A, mid+1, right)
		Merge(A, left, mid, right)
	}
}
