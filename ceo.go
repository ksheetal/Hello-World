package main

import "fmt"

type CEO struct {
	g  *GreetService
	f  *Finance
	i  *Intern
	hr *HR
}

func NewCEO(g *GreetService, f *Finance, hr *HR) *CEO {
	return &CEO{
		g:  g,
		f:  f,
		hr: hr,
	}
}

func (c *CEO) Greet() {
	c.g.Greet("Hello From CEO")
}

func (c *CEO) AskDoubts() {
	c.g.Greet("Who hired you all?")
}

func (c *CEO) FireAll() {
	fmt.Println("firing all")
}
