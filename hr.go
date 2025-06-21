package main

type HR struct {
	g *GreetService
}

func NewHR(g *GreetService) *HR {
	return &HR{
		g: g,
	}
}

func (hr *HR) Greet() {
	hr.g.Greet("Hello From HR")
}

func (hr *HR) GivePromotion() {
	hr.g.Greet("Giving Promotion ")
}
