package main

import (
	"fmt"
	"strconv"
)

func main()  {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseInt("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("159")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}
