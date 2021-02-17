package widgets

import (
	"github.com/liamg/flinch/components"
	"github.com/liamg/flinch/core"
	"github.com/liamg/flinch/window"
)

func Confirm(msg string) (bool, error) {

	win, err := window.New()
	if err != nil {
		return false, err
	}

	minLength := 50
	maxLength := minLength
	if len(msg) > maxLength {
		maxLength = len(msg)
	}

	minSize := core.SizeStrategyMultiple(
		core.SizeStrategyPercentage(80, 0),
		core.SizeStrategyAtLeast(core.Size{W: minLength, H: 1}),
		core.SizeStrategyAtMost(core.Size{W: maxLength + 8, H: 100}),
	)

	text := components.NewText(msg)
	text.SetSizeStrategy(minSize)
	textFrame := components.NewFrame(text)

	strip := components.NewColumnLayout()
	strip.SetSizeStrategy(minSize)

	yes := components.NewButton("Yes")
	no := components.NewButton("No")
	strip.Add(yes)
	strip.Add(no)

	rows := components.NewRowLayout()
	rows.Add(textFrame)
	rows.Add(components.NewSpacer(core.Size{H: 1}))
	rows.Add(strip)
	rows.SetAlignment(core.AlignCenter)

	win.SetAlignment(core.AlignCenter)
	win.Add(rows)

	var confirmed bool

	yes.OnPress(func() {
		confirmed = true
		win.Close()
	})

	no.OnPress(func() {
		win.Close()
	})

	if err := win.Show(); err != nil {
		return false, err
	}

	return confirmed, nil
}
