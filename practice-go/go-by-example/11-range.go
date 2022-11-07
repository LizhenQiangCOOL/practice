package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	// 遍历数组：index， value
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum：", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//遍历map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("keys:", k)
	}

	//遍历字符川， 其中value为 ASCII
	for i, c := range "go" {
		fmt.Println(i, c)
		fmt.Printf("string(c): %v \n", string(c))
	}
}
