package facade

import "fmt"

// API 这个包待实现的接口
type API interface {
	Test() string
}

func NewAPI() API {
	return &apiImpl{
		a: NewAMoudleAPI(),
		b: NewBModuleAPI(),
	}
}

// apiImpl 对外接口的
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s \n %s", aRet, bRet)
}

// AModeleAPI 子系统A的接口
type AModuleAPI interface {
	TestA() string
}
type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

func NewAMoudleAPI() AModuleAPI {
	return &aModuleImpl{}
}

// BModuleAPI 子系统B的接口
type BModuleAPI interface {
	TestB() string
}
type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}
