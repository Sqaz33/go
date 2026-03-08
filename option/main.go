package main

import "fmt"

type Point struct {
	X, Y int
}

type Option func(*Point)

func SetXY(x, y int) Option {
	return func(p *Point) {
		p.X = x
		p.Y = y
	}
}

func NewPoint(opts ...Option) (p Point) {
	p = Point{1, 2}

	for _, o := range opts {
		o(&p)
	}
	return
}

func main() {
	p := NewPoint(SetXY(10, 10))
	fmt.Println(p)
}
