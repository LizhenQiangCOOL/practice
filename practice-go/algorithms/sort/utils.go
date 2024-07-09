package sort

import (
	"math/rand"
	"time"
)

func RandomSlice(n int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	a := make([]int, n)
	for i := 0; i < len(a); i++ {
		a[i] = r.Intn(len(a))
	}
	return a
}

// CheckIsSorted 检查是否从小到大排序
func CheckIsSorted(s []int) bool {
	for i, _ := range s {
		if i+1 < len(s) {
			if s[i] > s[i+1] {
				return false
			}
		}
	}
	return true
}
