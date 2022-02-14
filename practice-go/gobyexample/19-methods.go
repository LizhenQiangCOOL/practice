package main

import "fmt"

type rect struct{
	width, height int
}

// 给结构体绑定方法， 起面的 r *rect 叫接收器
// 如果接收器是指针，则可以在方法内部修改结构体
func (r *rect) area() int{
	return r.width * r.height
}

func (r rect) perim() int{
	return 2*r.width + 2*r.height
}

func main()  {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim", r.perim())

	r = rect{width: 10, height: 5}
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim", rp.perim())
}
