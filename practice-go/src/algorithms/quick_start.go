package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
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

	//copy 删除a[i] 可以用copy将i+1到末尾的值覆盖到i，然后末尾-1
	a := make([]int, 0)
	var i int = 0
	copy(a[i:], a[i+1:])
	a = a[:len(a)-1]
	// make 创建长度，则通过索引赋值
	a = make([]int, 5)
	a[1] = 1
	// make 创建长度为0， 则通过append()赋值
	a = make([]int, 0)
	a = append(a, 1)

	// 常用技巧
	//类型转换
	str := "12345"                         // s[0] 类型是byte
	num := int(str[0] - '0')               //1
	str1 := string(str[0])                 // "i"
	b := byte(num + '0')                   // 'i'
	fmt.Printf("%d %s %c\n", num, str1, b) //111
	//字符串转换成数字
	strnew := "1"
	num, _ = strconv.Atoi(strnew) //
	strnew = strconv.Itoa(num)
	fmt.Println(num, strnew)
}
