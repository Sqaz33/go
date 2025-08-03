// a package is a way to group functions, and it's made up of all the files in the same directory
package main

// go mod init example/hello
// go build .
// go run .
// go mod tidy or o get rsc.io/quote

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
