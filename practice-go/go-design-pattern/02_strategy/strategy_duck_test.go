package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMallarDuck(t *testing.T) {
	t.Run("test mallarDuck", func(t *testing.T) {
		// 能飞 会叫 的鸭子
		q := Quack{}
		f := FlyWithWings{}
		d := Duck{}
		mallarDuck, err := NewMallarDuck(d, f, q)
		assert.NoError(t, err)
		mallarDuck.display()
		mallarDuck.Ducker.swim()
		mallarDuck.FlyBehavior.fly()
		mallarDuck.QuackBehavior.quack()
	})
}
