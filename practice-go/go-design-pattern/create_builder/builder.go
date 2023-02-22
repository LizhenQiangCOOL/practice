package builder

type Product struct {
	Part1 string
	Part2 string
	Part3 string
}

type Builder interface {
	Part1()
	Part2()
	Part3()
	GetResult() Product
}

type ConcreteBuilderA struct {
	result Product
}

func (b *ConcreteBuilderA) Part1() {
	b.result.Part1 = "part1"
}

func (b *ConcreteBuilderA) Part2() {
	b.result.Part2 = "part2"
}
func (b *ConcreteBuilderA) Part3() {
	b.result.Part3 = "part3"
}

func (b *ConcreteBuilderA) GetResult() Product {
	return b.result
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}
