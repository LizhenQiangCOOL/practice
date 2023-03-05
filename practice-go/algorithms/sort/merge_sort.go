package sort

import "math"

func Merge(A []int, p int, q int, r int) {
	// n1 第一个数组长度，n2 第二个数组长度
	n1 := q - p + 1
	n2 := r - q //满足r-q为第二个数组长度即可
	//let L[0..n1] and R[0...n2] be new arrays
	L := make([]int, n1+1)
	R := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		L[i] = A[p+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[q+i+1]
	}
	// 增加哨兵
	L[n1] = math.MaxInt
	R[n2] = math.MaxInt
	i, j := 0, 0
	for k := p; k <= r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i += 1
		} else {
			A[k] = R[j]
			j += 1
		}
	}
}

// A[p..r]
func MergeSort(A []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(A, p, q)
		MergeSort(A, q+1, r)
		Merge(A, p, q, r)
	}
}
