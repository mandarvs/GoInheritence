package main

import "fmt"

type greeter struct {
	greeting string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting)
}

func main() {
	g := greeter{greeting: "Hello, world!"}
	g.greet()
}
