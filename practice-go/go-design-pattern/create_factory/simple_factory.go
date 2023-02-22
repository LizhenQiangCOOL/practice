package factory

type Operation interface {
	SetA(float64)
	SetB(float64)
	Result() float64
}

type Addition struct {
	a, b float64
}

func (add *Addition) SetA(a float64) {
	add.a = a
}

func (add *Addition) SetB(b float64) {
	add.b = b
}

func (add *Addition) Result() float64 {
	return add.a + add.b
}

type Subtraction struct {
	a, b float64
}

func (sub *Subtraction) SetA(a float64) {
	sub.a = a
}

func (sub *Subtraction) SetB(b float64) {
	sub.b = b
}

func (sub *Subtraction) Result() float64 {
	return sub.a - sub.b
}

type SimpleFactory struct{}

// 根据传入的条件，返回不同的结构体
func (f *SimpleFactory) NewOperation(operator string) Operation {
	switch operator {
	case "+":
		return &Addition{}
	case "-":
		return &Subtraction{}
	default:
		panic("unsupported operator")
	}
}
