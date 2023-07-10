package main

import "fmt"

type Simpler interface { // interface implementing functions called on struct Simple
	Get() int
	Set(n int)
}

type Simple struct {
	n int
}

func (p *Simple) Get() int {
	return p.n
}

func (p *Simple) Set(u int) {
	p.n = u
	return
}

func fI(it Simpler) int { // function calling both methods through interface
	it.Set(5)
	return it.Get()
}

func main() {
	s := &Simple{}
	fmt.Println(fI(Simpler(s)))
}
