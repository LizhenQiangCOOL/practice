package factory

// Operator 是被封装的实际类接口 [抽象产品类]
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口 [定义抽象工厂类]
type OperatorFactory interface {
	Create() Operator
}

// PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{}
}

// PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	a int
	b int
}

func (o *PlusOperator) SetA(number int) {
	o.a = number
}
func (o *PlusOperator) SetB(number int) {
	o.b = number
}

// Result 获取结果
func (o *PlusOperator) Result() int {
	return o.a + o.b
}

// MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{}
}

// MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	a int
	b int
}

func (o *MinusOperator) SetA(number int) {
	o.a = number
}
func (o *MinusOperator) SetB(number int) {
	o.b = number
}

// Result 获取结果
func (o *MinusOperator) Result() int {
	return o.a - o.b
}
