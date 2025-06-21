package main

type Intern struct {
	g *GreetService
}

func NewIntern(g *GreetService) *Intern {
	return &Intern{
		g: g,
	}
}

func (i *Intern) Greet() {
	i.g.Greet("Hello From HR")
}

func (i *HR) AskDoubts() {
	i.g.Greet("What is your total experience?")
}
