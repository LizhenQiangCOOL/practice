package main

import "fmt"

func main()  {
	var a = "initial"
	fmt.Println(a)

	//多个变量一起赋值
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//默认赋值 0, nil, "", [], {}
	var e int
	fmt.Println(e)


	//短声明
	f := "apple"
	fmt.Println(f)
}
