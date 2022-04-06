package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	// 栈 （后进先出）
	stack := make([]int, 0)
	stack = append(stack, 10)
	v1 := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(len(stack) == 0)
	fmt.Println(v1)

	// 队列 (先进先出)
	queue := make([]int, 0)
	queue = append(queue, 10)
	v2 := queue[0]
	queue = queue[1:]
	fmt.Println(len(queue) == 0)
	fmt.Println(v2)

	// 字典 (hash表)
	m := make(map[string]int)
	m["hello"] = 1
	delete(m, "hello")
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 标准库sort
	sort.Ints([]int{})
	sort.Strings([]string{})
	s := make([]int, 0)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	// math
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)
}
