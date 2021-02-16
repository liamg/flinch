package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/components"
	"github.com/liamg/flinch/core"
	"github.com/liamg/flinch/window"
)

func PasswordInput(msg string) (string, error) {

	win, err := window.New()
	if err != nil {
		return "", err
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

	inputbox := components.NewPasswordInput()
	inputbox.SetSizeStrategy(minSize)
	listFrame := components.NewFrame(inputbox)

	text := components.NewText(msg)
	text.SetSizeStrategy(minSize)
	textFrame := components.NewFrame(text)

	buttons := components.NewColumnLayout()
	buttons.SetSizeStrategy(minSize)

	var selected bool

	buttons.Add(components.NewSpacer(core.Size{W: 1}))

	okButton := components.NewButton("OK")
	okButton.OnPress(func() {
		selected = true
		win.Close()
	})
	buttons.Add(okButton)

	inputbox.OnKeypress(func(key *tcell.EventKey) bool {
		switch key.Key() {
		case tcell.KeyEnter:
			selected = true
			win.Close()
			return true
		}
		return false
	})

	buttons.Add(components.NewSpacer(core.Size{W: 1}))

	cancelButton := components.NewButton("Cancel")
	cancelButton.OnPress(func() {
		win.Close()
	})
	buttons.Add(cancelButton)

	help := components.NewText("Use TAB, ENTER")
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
		return "", err
	}

	if !selected {
		return "", ErrInputCancelled
	}

	input := inputbox.GetInput()

	return input, nil
}
