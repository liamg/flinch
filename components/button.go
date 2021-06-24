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

func (l *button) Text() string {
	return l.label
}

func (l *button) Render(canvas core.Canvas) {

	st := core.StyleButton
	if l.selected {
		st = core.StyleButtonSelected
	}

	canvas.Fill(' ', st)
	size := canvas.Size()

	if l.selected {
		edges := st.GetButtonEdges()
		canvas.Set(0, 0, edges[0], st.ToggleInvert())
		canvas.Set(size.W-1, 0, edges[1], st.ToggleInvert())
	}

	for i := 0; i < len(l.label); i++ {
		canvas.Set(((size.W-len(l.label))/2)+i, 0, rune(l.label[i]), st)
	}
}

func (l *button) MinimumSize() core.Size {
	return core.Size{W: len(l.label) + 4, H: 1}
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

func (l *button) Selected() bool {
	return l.selected
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
