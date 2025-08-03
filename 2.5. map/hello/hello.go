package main

// go mod edit -replace example.com/greetings=../greetings
import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Stepan1", "Stepan2", "Stepan3"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range messages {
		fmt.Println(msg)
	}
}
