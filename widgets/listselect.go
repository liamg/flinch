package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/components"
	"github.com/liamg/flinch/core"
	"github.com/liamg/flinch/window"
)

func ListSelect(msg string, options []string) (int, string, error) {

	win, err := window.New()
	if err != nil {
		return -1, "", err
	}

	minLength := 50
	maxLength := minLength
	if len(msg) > maxLength {
		maxLength = len(msg)
	}
	for _, opt := range options {
		if len(opt) > maxLength {
			maxLength = len(opt)
		}
	}

	minSize := core.SizeStrategyMultiple(
		core.SizeStrategyPercentage(80, 0),
		core.SizeStrategyAtLeast(core.Size{W: minLength, H: 1}),
		core.SizeStrategyAtMost(core.Size{W: maxLength + 8, H: 100}),
	)

	list := components.NewListSelect(options)
	list.SetSizeStrategy(minSize)
	listFrame := components.NewFrame(list)

	text := components.NewText(msg)
	text.SetSizeStrategy(minSize)
	textFrame := components.NewFrame(text)

	strip := components.NewColumnLayout()
	strip.SetSizeStrategy(minSize)

	var selected bool

	list.OnKeypress(func(key *tcell.EventKey) bool {
		switch key.Key() {
		case tcell.KeyEnter:
			selected = true
			win.Close()
			return true
		}
		return false
	})

	help := components.NewText("ENTER to confirm, ESC to cancel")
	help.SetSizeStrategy(core.SizeStrategyMaximumWidth())
	help.SetAlignment(core.AlignCenter)
	help.SetStyle(core.StyleFaint)
	strip.Add(help)

	rows := components.NewRowLayout()
	rows.Add(textFrame)
	rows.Add(listFrame)
	rows.Add(components.NewSpacer(core.Size{H: 1}))
	rows.Add(strip)
	rows.SetAlignment(core.AlignCenter)

	win.SetAlignment(core.AlignCenter)
	win.Add(rows)

	if err := win.Show(); err != nil {
		return -1, "", err
	}

	index, str := list.GetSelection()

	if !selected || index == -1 {
		return -1, "", ErrInputCancelled
	}

	return index, str, nil
}
