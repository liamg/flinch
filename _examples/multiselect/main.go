package main

import (
	"fmt"
	"strings"

	"github.com/liamg/flinch/widgets"
)

func main() {

	_, items, err := widgets.MultiSelect(
		"Select one or more environment(s)...",
		[]string{
			"Development",
			"Test",
			"Staging",
			"Production",
		},
	)
	if err == widgets.ErrInputCancelled {
		fmt.Println("User cancelled.")
		return
	}

	fmt.Printf("You selected %s.\n", strings.Join(items, ", "))
}
