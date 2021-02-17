package main

import (
	"fmt"

	"github.com/liamg/flinch/widgets"
)

func main() {

	confirmed, err := widgets.Confirm("Are you sure?")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Confirmed: %t\n", confirmed)
}
