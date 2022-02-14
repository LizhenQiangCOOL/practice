package main

import "fmt"

type person struct {
	name string
	age int
}

// 返回结构体指针
func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main()  {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age:30})

	fmt.Println(person{name: "Fred"})

	// 结构体指针
	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))


	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}