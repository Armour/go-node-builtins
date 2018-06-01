package main

import (
	"fmt"

	"github.com/Armour/go-node-builtins/builtins"
)

func main() {
	b, err := builtins.GetVersion("6.0.0")
	if err != nil {
		// handle error
	}
	fmt.Printf("%v", b)
}
