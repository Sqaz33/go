package main // точка входа всегда в package main

import (
	"example/hello/internal/add"
	"fmt"
)

func main() {
	fmt.Println(add.Add(1, 2))
}
