package main

import "fmt"

func main() {
	// 8 字节 byte = 64位 bit
	var intArr [3]int64

	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	// 1 byte = 8 bit
	fmt.Printf("intArr的地址=%p int[0]地址=%p int[1]地址%p int[2]地址%p", &intArr, &intArr[0], &intArr[1], &intArr[2])
	fmt.Println(intArr)
}
