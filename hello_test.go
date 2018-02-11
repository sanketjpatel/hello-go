package main

import "testing"

func TestSayHello(t *testing.T) {
	expected := "hello, world\n"
	result := SayHelloWorld()
	if result != expected {
		t.Errorf("SayHello() == %q, expected == %q", result, expected)
	}
}
