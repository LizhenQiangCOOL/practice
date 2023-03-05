package sort

import (
	"math/rand"
	"time"
)

func RandomSlice(n int) []int {
	rand.Seed(time.Now().Unix())
	r := make([]int, n)
	for i := 0; i < len(r); i++ {
		r[i] = rand.Intn(len(r))
	}
	return r
}

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
