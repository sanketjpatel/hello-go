package main

import "fmt"

// sayHelloWorld returns a string "hello, world\n"
func sayHelloWorld() string {
	return "Hello, world\n"
}

func main() {
	fmt.Print(sayHelloWorld())
}
