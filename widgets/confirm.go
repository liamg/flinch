package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/components"
	"github.com/liamg/flinch/core"
	"github.com/liamg/flinch/window"
)

// Confirm displays a yes/no dialog with the given message. If "Yes" is selected, true is returned.
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
	text.SetAlignment(core.AlignCenter)
	text.SetSizeStrategy(minSize)

	strip := components.NewColumnLayout()
	strip.SetSizeStrategy(minSize)
	strip.SetAlignment(core.AlignCenter)

	yes := components.NewButton("Yes")
	no := components.NewButton("No")
	strip.Add(yes)
	strip.Add(no)

	rows := components.NewRowLayout()
	rows.Add(text)
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

	win.OnKeypress(func(key *tcell.EventKey) bool {
		switch key.Key() {
		case tcell.KeyLeft:
			strip.Deselect()
			strip.Select()
		case tcell.KeyRight:
			strip.Select()
		case tcell.KeyRune:
			switch key.Rune() {
			case 'y', 'Y':
				confirmed = true
				win.Close()
			case 'n', 'N':
				confirmed = false
				win.Close()
			}
		}
		return false
	})

	if err := win.Show(); err != nil {
		return false, err
	}

	return confirmed, nil
}
