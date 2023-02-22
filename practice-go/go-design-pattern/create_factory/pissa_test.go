package factory

import (
	"testing"
)

func Test_NewNYStyleStore(t *testing.T) {
	nyStore := NewNYStyleStore()
	pissa := nyStore.orderPizza("cheese")
	if pissa.GetName() != "cheese" {
		t.Fatal("pisss order error")
	}
}
