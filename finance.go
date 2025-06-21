package main

type Finance struct {
	g *GreetService
}

func NewFinance(g *GreetService) *Finance {
	return &Finance{
		g: g,
	}
}

func (f *Finance) Greet() {
	f.g.Greet("Hello From Finance")
}

func (f *Finance) AskDoubts() {
	f.g.Greet("What is the salary of intern?")
}
