package main

func main() {

	hr := &HR{g: &GreetService{}}
	f := &Finance{g: &GreetService{}}
	i := &Intern{g: &GreetService{}}

	c := &CEO{g: &GreetService{}, f: f, hr: hr, i: i}

	//hr.Greet()
	//hr.AskDoubts()
	//
	//f.Greet()
	//f.AskDoubts()

	c.Greet()
	c.AskDoubts()
	c.FireAll()
}
