package main

import (
	"github.com/liamg/flinch/components"
	"github.com/liamg/flinch/core"
	window2 "github.com/liamg/flinch/window"
)

func main() {

	window, err := window2.New()
	if err != nil {
		panic(err)
	}

	textA := components.NewText("AAAAA").SetAlignment(core.AlignCenter)
	textA.SetSizeStrategy(core.SizeStrategyMaximumWidth())
	textB := components.NewText("BBBBB").SetAlignment(core.AlignCenter)
	textB.SetSizeStrategy(core.SizeStrategyMaximumWidth())
	textC := components.NewText("CCCCC").SetAlignment(core.AlignCenter)
	textC.SetSizeStrategy(core.SizeStrategyMaximumWidth())

	//window.SetAlignment(core.AlignCenter)
	window.Add(textA)
	window.Add(textB)
	window.Add(textC)

	if err := window.Show(); err != nil {
		panic(err)
	}
}
