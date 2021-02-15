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

	list := components.NewMultiListSelect([]string{"A", "bcd", "EF", "gGgGgGgG"})

	limiter := components.NewLimiter(list)
	limiter.WithPercentageLimitOnWidth(60)

	text := components.NewText("Select your whatever...")
	textFrame := components.NewFrame(text)
	textLimiter := components.NewLimiter(textFrame)
	textLimiter.WithPercentageLimitOnWidth(60)

	buttons := components.NewColumnLayout()
	buttons.Add(components.NewButton("OK"))
	buttons.Add(components.NewButton("Cancel"))

	rows := components.NewRowLayout()
	rows.Add(textLimiter)
	rows.Add(limiter)
	rows.Add(buttons)
	rows.WithJustification(core.JustifyCenter)

	window.WithJustification(core.JustifyCenter)
	window.Add(rows)

	if err := window.Show(); err != nil {
		panic(err)
	}
}
