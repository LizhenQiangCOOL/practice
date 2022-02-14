package main

import "fmt"

func main()  {

	// map使用make创建不需要长度
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	//删除 key
	delete(m, "k2")
	fmt.Println("map:", m)

	//安全取出不存在的key, prs为false
	_, prs := m["k2"]
	fmt.Println("prs:", prs)


	n := map[string]int{"foo": 1, "bar":2}
	fmt.Println("map:", n)
}
