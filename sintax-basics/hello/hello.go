package main

import "fmt"

func main() {
	sayHello("Go")
}

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
