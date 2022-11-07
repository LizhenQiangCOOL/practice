package main

import "fmt"

// closurse   闭包
// 个人认为： 函数内嵌函数，而内函数使用了外函数的变量，跟外函数的变量绑定了，该变量像有记忆一样
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

}
