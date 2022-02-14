package main

import "fmt"

// 返回多值
func vals() (int, int){
	return 3, 7
}

func main(){

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
