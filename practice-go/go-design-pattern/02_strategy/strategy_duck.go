package strategy

import "fmt"

// 策略模式 实现鸭子类
// 鸭子类： display 怎么样的鸭子， quack 会叫，fly 会飞
type Ducker interface {
	swim()
}

// 鸭子的通用方法
type Duck struct {
}

func (d Duck) swim() {
	fmt.Println("all duck float, even decoys!")
}

// 各种飞行行为
type FlyBehavior interface {
	fly()
}

type FlyWithWings struct {
}

func (f FlyWithWings) fly() {
	//实现鸭子飞行
	fmt.Println("I can fly")
}

type FlyNoWay struct {
}

func (f FlyNoWay) fly() {
	// 什么都不做，笨鸭子不会飞
	fmt.Println("I can not fly  hahaha")
}

// 各种喊叫行为
type QuackBehavior interface {
	quack()
}

type Quack struct {
}

func (q Quack) quack() {
	//呱呱叫
	fmt.Println("ga ga ga ga ga quack")
}

type Squeak struct {
}

func (q Squeak) quack() {
	//吱吱叫
	fmt.Println("zhi zhi zhi  zhi zhi quack")
}

type MuteQuack struct {
}

func (q MuteQuack) quack() {
	//又是一只笨鸭，不会叫
	fmt.Println("I can not quack sorry")
}

// 组合代替继承(结构体内嵌)
type MallarDuck struct {
	Ducker        Ducker
	FlyBehavior   FlyBehavior
	QuackBehavior QuackBehavior
}

func NewMallarDuck(d Ducker, f FlyBehavior, q QuackBehavior) (*MallarDuck, error) {
	return &MallarDuck{
		Ducker:        d,
		FlyBehavior:   f,
		QuackBehavior: q,
	}, nil
}

func (m MallarDuck) display() {
	fmt.Println("I'm a real Mallard duck")
}
