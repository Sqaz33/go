package main

// go mod edit -replace example.com/greetings=../greetings
import (
	"fmt"

	"example.com/greetings"
)

func main() {
	var msg string = greetings.Hello("Stepan")
	fmt.Println(msg)
}
