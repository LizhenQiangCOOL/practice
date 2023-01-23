package factory

import (
	"fmt"
)

type Pissa interface {
	GetName() string
	prepare()
	bake()
	cute()
	box()
}

type PissaBase struct {
	Name string
}

func (p *PissaBase) GetName() string {
	return p.Name
}

func (p *PissaBase) prepare() {
	fmt.Sprintf("%s pissa is prepared \n", p.Name)
}
func (p *PissaBase) bake() {
	fmt.Printf("%s pissa is bake \n", p.Name)
}
func (p *PissaBase) cute() {
	fmt.Printf("%s pissa is cute \n", p.Name)
}
func (p *PissaBase) box() {
	fmt.Printf("%s pissa is box \n", p.Name)
}

// 一种pissa
type CheesePissa struct {
	*PissaBase
}

func NewCheesePissa() *CheesePissa {
	return &CheesePissa{&PissaBase{Name: "cheese"}}
}

// 蔬菜pissa
type VeggiePissa struct {
	*PissaBase
}

func NewVeggiePissa() *VeggiePissa {
	return &VeggiePissa{&PissaBase{Name: "veggie"}}
}

type PissaStore interface {
	createPizza(t string) Pissa
	orderPizza(t string) Pissa
}

// PissaStore的一种实现
type NYStyleStore struct {
}

func NewNYStyleStore() *NYStyleStore {
	return &NYStyleStore{}
}

func (s *NYStyleStore) createPizza(t string) Pissa {
	var pissa Pissa
	if t == "cheese" {
		pissa = NewCheesePissa()
	} else if t == "veggie" {
		pissa = NewVeggiePissa()
	}
	return pissa
}
func (s *NYStyleStore) orderPizza(t string) Pissa {
	p := s.createPizza(t)
	// 可自定义其中一些操作
	p.prepare()
	p.bake()
	p.cute()
	p.box()
	return p
}

type ChicagoStyleStore struct{}

func NewChicagoStyleStore() *ChicagoStyleStore {
	return &ChicagoStyleStore{}
}
func (s *ChicagoStyleStore) createPizza() {

}
func (s *ChicagoStyleStore) orderPizza() {}
