package main

import "fmt"

type Speacker interface {
	Speack() string
}

type Dog struct{}

func (*Dog) Speack() string {
	return "Woof"
}

type Cat struct{}

func (Cat) Speack() string {
	return "Meow"
}

func MakeSound(s Speacker) {
	fmt.Println(s.Speack())
}

func main() {
	d := &Dog{}
	MakeSound(d)

	c := Cat{}
	MakeSound(c)
}
