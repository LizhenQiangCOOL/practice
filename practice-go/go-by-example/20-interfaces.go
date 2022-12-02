package main

import (
	"fmt"
	"math"
)

// geometry 接口等待被实现
// 实现了该接口的所有属性 =  实现了该接口
type geometry interface {
	area() float64
	perim() float64
}

type rectf struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectf) area() float64 {
	return r.width * r.height
}

func (r rectf) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 参数为 geometry
// 只要实现了接口内的方法，就可以视为实现了该接口
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rectf{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

}
