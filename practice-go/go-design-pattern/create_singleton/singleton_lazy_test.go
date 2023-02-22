package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	tests := []struct {
		name string
		want *Singleton
	}{
		{
			"懒汉式单例测试",
			GetInstance(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetLazyInstance(), "GetLazyInstance()")
		})
	}
}

func BenchmarkGetLazyInstance(b *testing.B) {
	//并行测试
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
