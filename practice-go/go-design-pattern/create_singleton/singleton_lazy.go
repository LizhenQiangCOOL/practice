package singleton

import "sync"

var (
	lazySingleton *Singleton
	once          sync.Once
)

// 懒汉式单例，可多 go 协程使用
func GetLazyInstance() *Singleton {
	once.Do(func() {
		// sync.Once 内部使用 done 标记实现双重锁定
		// done 标记为啥要放在结构体第一位： https://stackoverflow.com/questions/59174176/what-does-hot-path-mean-in-the-context-of-sync-once
		lazySingleton = &Singleton{}
	})
	return lazySingleton
}
