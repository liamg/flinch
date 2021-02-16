package main

import (
	"fmt"

	"github.com/liamg/flinch/widgets"
)

func main() {

	_, item, err := widgets.ListSelect(
		"Select an environment...",
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

	fmt.Printf("You selected %s.\n", item)
}
