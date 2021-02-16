package main

import (
	"fmt"

	"github.com/liamg/flinch/widgets"
)

func main() {

	name, err := widgets.Input("Enter your name...")
	if err == widgets.ErrInputCancelled {
		fmt.Println("User cancelled.")
		return
	}

	fmt.Printf("Hello, %s!\n", name)
}
