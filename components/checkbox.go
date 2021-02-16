package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type checkbox struct {
	label    string
	checked  bool
	selected bool
}

func NewCheckbox(label string, checked bool) *checkbox {
	return &checkbox{
		label:   label,
		checked: checked,
	}
}

func (t *checkbox) Text() string {
	return t.label
}

func (t *checkbox) SetChecked(checked bool) {
	t.checked = checked
}

func (t *checkbox) Render(canvas core.Canvas) {

	st := core.StyleDefault
	if t.selected {
		st = core.StyleSelected
	}

	canvas.Set(0, 0, '[', st)
	canvas.Set(2, 0, ']', st)

	if t.checked {
		canvas.Set(1, 0, '✔', st)
	} else {
		canvas.Set(1, 0, '_', st)
	}

	for i := 0; i < len(t.label); i++ {
		canvas.Set(4+i, 0, rune(t.label[i]), st)
	}
}

func (t *checkbox) Size(parent core.Canvas) (int, int) {
	w, _ := parent.Size()
	return w, 1
}

func (l *checkbox) Deselect() {
	l.selected = false
}

func (l *checkbox) Select(loop bool) bool {
	if l.selected {
		return false
	}
	l.selected = true
	return true
}

func (l *checkbox) HandleKeypress(key *tcell.EventKey) {
	switch key.Key() {
	case tcell.KeyRune:
		switch key.Rune() {
		case ' ':
			l.checked = !l.checked
		}
	}
}
