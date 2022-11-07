package singleton

// 单例模式  饿汉式单例
type Singleton struct{}

var singleton *Singleton

// 利用 init 初始化机制
func init() {
	singleton = &Singleton{}
}

// 获取实例
func GetInstance() *Singleton {
	return singleton
}
