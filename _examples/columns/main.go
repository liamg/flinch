package main

import (
	"github.com/liamg/flinch"
	"github.com/liamg/flinch/components"
	"github.com/liamg/flinch/core"
)

func main() {

	window, err := flinch.New()
	if err != nil {
		panic(err)
	}

	textA := components.NewText("A")
	textA.SetJustification(core.JustifyCenter)
	textB := components.NewText("B")
	textB.SetJustification(core.JustifyCenter)
	textC := components.NewText("C")
	textC.SetJustification(core.JustifyCenter)

	window.SetJustification(core.JustifyFill)
	window.Add(textA)
	window.Add(textB)
	window.Add(textC)

	if err := window.Show(); err != nil {
		panic(err)
	}
}
