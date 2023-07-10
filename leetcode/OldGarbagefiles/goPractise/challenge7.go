package main

import "fmt"

type Square struct {
	side float32
}

type Triangle struct { // implement this struct
	base   float32
	height float32
}

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface { // implement this interface
	Perimeter() float32
}

func main() {
	s := &Square{2}
	t := &Triangle{2, 3}
	var areaInf AreaInterface
	var periInf PeriInterface

	areaInf = s
	fmt.Println(areaInf.Area())
	periInf = s
	fmt.Println(periInf.Perimeter())

	areaInf = t
	fmt.Println(areaInf.Area())
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 { // implement method called on square to calculate its perimeter
	return sq.side * sq.side * 2
}

func (tr *Triangle) Area() float32 { // implement method called on triangle to calculate its area
	return 0.5 * tr.base * tr.height
}
