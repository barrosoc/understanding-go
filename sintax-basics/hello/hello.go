package main

import "fmt"

func main() {
	sayHello("Go")
}

func sayHello(s string) {
	fmt.Printf("Hello, %s!\n", s)
}
