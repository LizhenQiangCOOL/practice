package main

import "fmt"

func f(slice []int) []int {
	if len(slice) > 2 {
		slice[2] = 100
	}
	return slice
}

func main() {
	s := make([]int, 3, 5)
	s[0], s[1], s[2] = 0, 1, 2
	s1 := f(s)
	fmt.Printf("s: %+v \n", s)
	fmt.Printf("s1:%+v \n", s1)
}
