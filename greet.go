package main

import "fmt"

type GreetService struct {
}

func (g *GreetService) Greet(str string) {
	fmt.Println(str)
}

type Greeter interface {
	Greet(str string)
}
