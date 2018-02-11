package main

import "fmt"

// SayHelloWorld returns a string "hello, world\n"
func SayHelloWorld() string {
	return "hello, world\n"
}

func main() {
	fmt.Print(SayHelloWorld())
}
