package prototype

// Cloneeable 是原形对象需要实现的接口
type Cloneable interface {
	//遇到 slice 或map 考虑要不要深拷贝
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{prototypes: make(map[string]Cloneable)}
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name].Clone()
}
func (p *PrototypeManager) Set(name string, value Cloneable) {
	p.prototypes[name] = value
}
