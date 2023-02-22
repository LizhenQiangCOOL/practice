package builder

import "fmt"

func ExampleBuilder() {
	builder := &ConcreteBuilderA{}
	director := NewDirector(builder)
	director.Construct()
	product := builder.GetResult()
	fmt.Println(product)
	// Output:
	// {part1 part2 part3}
}
