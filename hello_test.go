package main

import "testing"

func TestSayHello(t *testing.T) {
	expected := "Hello, world\n"
	result := sayHelloWorld()
	if result != expected {
		t.Errorf("SayHello() == %q, expected == %q", result, expected)
	}
}
