package factory

import "testing"

func Test_hiAPI_Say(t *testing.T) {
	api := NewAPI("hi")
	s := api.Say("Tom")
	if s != "Hi Tom" {
		t.Fatal("Type1 test fail")
	}
}

func Test_helloAPI_Say(t *testing.T) {
	api := NewAPI("hello")
	s := api.Say("Tom")
	if s != "Hello Tom" {
		t.Fatal("Type1 test fail")
	}
}
