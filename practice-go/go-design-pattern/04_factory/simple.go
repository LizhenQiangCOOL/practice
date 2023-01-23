package factory

import "fmt"

// 简单工厂

type API interface {
	Say(name string) string
}

// 返回不同的API实例
func NewAPI(t string) API {
	if t == "hi" {
		return &hiAPI{}
	} else if t == "hello" {
		return &helloAPI{}
	}
	return nil
}

// hiApi API的一种实现
type hiAPI struct{}

func (h *hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi %s", name)
}

// helloApi API的另外一种实现
type helloAPI struct{}

func (h *helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
