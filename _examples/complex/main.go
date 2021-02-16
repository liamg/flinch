package main

import (
	"fmt"
	"strings"

	"github.com/liamg/flinch/components"
	"github.com/liamg/flinch/core"
	"github.com/liamg/flinch/window"
)

func main() {

	win, err := window.New()
	if err != nil {
		panic(err)
	}

	minSize := core.SizeStrategyMultiple(
		core.SizeStrategyPercentage(80, 0),
		core.SizeStrategyAtLeast(core.Size{50, 1}),
		core.SizeStrategyAtMost(core.Size{60, 100}),
	)

	list := components.NewMultiListSelect([]string{"Development", "Test", "Production", "Management"})
	list.SetSizeStrategy(minSize)
	listFrame := components.NewFrame(list)

	text := components.NewText("Select one or more environment(s):")
	text.SetSizeStrategy(minSize)
	textFrame := components.NewFrame(text)

	buttons := components.NewColumnLayout()
	buttons.SetSizeStrategy(minSize)

	var selected bool

	buttons.Add(components.NewSpacer(core.Size{W: 1}))

	okButton := components.NewButton("OK")
	okButton.OnPress(func() {
		win.Close()
		selected = true
	})
	buttons.Add(okButton)

	buttons.Add(components.NewSpacer(core.Size{W: 1}))

	cancelButton := components.NewButton("Cancel")
	cancelButton.OnPress(func() {
		win.Close()
	})
	buttons.Add(cancelButton)

	help := components.NewText("Use UP/DOWN, TAB, ENTER")
	help.SetSizeStrategy(core.SizeStrategyMaximumWidth())
	help.SetAlignment(core.AlignRight)
	help.SetStyle(core.StyleFaint)
	buttons.Add(help)

	rows := components.NewRowLayout()
	rows.Add(textFrame)
	rows.Add(listFrame)
	rows.Add(components.NewSpacer(core.Size{H: 1}))
	rows.Add(buttons)
	rows.SetAlignment(core.AlignCenter)

	win.SetAlignment(core.AlignCenter)
	win.Add(rows)

	if err := win.Show(); err != nil {
		panic(err)
	}

	if selected {
		_, items := list.GetSelection()
		fmt.Printf("You selected %s.\n", strings.Join(items, ", "))
	} else {
		fmt.Println("User cancelled.")
	}
}
