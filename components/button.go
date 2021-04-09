package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type button struct {
	core.Sizer
	label     string
	selected  bool
	pressFunc func()
}

func NewButton(label string) *button {
	return &button{
		label: label,
	}
}

func (t *button) Text() string {
	return t.label
}

func (t *button) Render(canvas core.Canvas) {

	st := core.StyleButton
	if t.selected {
		st = core.StyleButtonSelected
	}

	canvas.Fill(' ', st)
	size := canvas.Size()

	if t.selected {
		canvas.Set(0, 0, '', st.ToggleInvert())
		canvas.Set(size.W-1, 0, '', st.ToggleInvert())
	}

	for i := 0; i < len(t.label); i++ {
		canvas.Set(((size.W-len(t.label))/2)+i, 0, rune(t.label[i]), st)
	}
}

func (t *button) MinimumSize() core.Size {
	return core.Size{W: len(t.label) + 4, H: 1}
}

func (l *button) Deselect() {
	l.selected = false
}

func (l *button) Select() bool {
	if l.selected {
		return false
	}
	l.selected = true
	return true
}

func (l *button) OnPress(f func()) {
	l.pressFunc = f
}

func (l *button) HandleKeypress(key *tcell.EventKey) {
	switch key.Key() {
	case tcell.KeyEnter:
		if l.pressFunc != nil {
			l.pressFunc()
		}
	case tcell.KeyRune:
		switch key.Rune() {
		case ' ':
			if l.pressFunc != nil {
				l.pressFunc()
			}
		}
	}
}
