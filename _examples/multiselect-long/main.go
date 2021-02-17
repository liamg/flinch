package main

import (
	"fmt"
	"strings"

	"github.com/liamg/flinch/widgets"
)

func main() {

	var options []string
	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	_, items, err := widgets.MultiSelect(
		"Select an option...",
		options,
	)
	if err == widgets.ErrInputCancelled {
		fmt.Println("User cancelled.")
		return
	}

	fmt.Printf("You selected %s.\n", strings.Join(items, ", "))
}
