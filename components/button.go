package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type button struct {
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

	st := core.StyleDefault
	if t.selected {
		st = core.StyleSelected
	}

	w, _ := canvas.Size()

	if t.selected {
		canvas.Set(1, 0, '>', st)
		canvas.Set(w-2, 0, '<', st)
	}

	for i := 0; i < len(t.label); i++ {
		canvas.Set(((w-len(t.label))/2)+i, 0, rune(t.label[i]), st)
	}
}

func (t *button) Size(parent core.Canvas) (int, int) {
	return len(t.label) + 6, 3
}

func (l *button) ToggleSelect(loop bool) bool {
	l.selected = !l.selected
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
