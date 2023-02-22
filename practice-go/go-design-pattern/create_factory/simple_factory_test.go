package factory

import "testing"

func TestSimpleFactory(t *testing.T) {
	f := &SimpleFactory{}

	add := f.NewOperation("+")
	add.SetA(1)
	add.SetB(2)
	result := add.Result()
	if result != 3 {
		t.Errorf("unexpected result: %f", result)
	}

	sub := f.NewOperation("-")
	sub.SetA(3)
	sub.SetB(1)
	result = sub.Result()
	if result != 2 {
		t.Errorf("unexpected result: %f", result)
	}
}
