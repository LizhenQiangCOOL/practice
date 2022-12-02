package observer

import "fmt"

// subject
type ISubject interface {
	Regiser(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

// 观察者
type IObserver interface {
	Update(msg string)
}

type Subject struct {
	observers []IObserver
}

// 注册
func (sub *Subject) Register(observer IObserver) {
	sub.observers = append(sub.observers, observer)
}

// 移除观察者
func (sub *Subject) Remove(observer IObserver) {
	// 可以使用map
	for i, ob := range sub.observers {
		if ob == observer {
			sub.observers = append(sub.observers[:i], sub.observers[i+1:]...)
		}
	}
}

// Notify 通知
func (sub *Subject) Notify(msg string) {
	for _, o := range sub.observers {
		o.Update(msg)
	}
}

// observer1
type Observer1 struct{}

// 实现观察者接口
func (o Observer1) Update(msg string) {
	fmt.Printf("Observer1: %s \n", msg)
}

// observer2
type Observer2 struct{}

// 实现观察者接口
func (o Observer2) Update(msg string) {
	fmt.Printf("Observer2: %s \n", msg)
}
