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

	textA := components.NewText("AAAAA").WithJustification(core.JustifyCenter)
	textB := components.NewText("BBBBB").WithJustification(core.JustifyCenter)
	textC := components.NewText("CCCCC").WithJustification(core.JustifyCenter)

	window.WithJustification(core.JustifyFill)
	window.Add(textA)
	window.Add(textB)
	window.Add(textC)

	if err := window.Show(); err != nil {
		panic(err)
	}
}
