package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type checkbox struct {
	label *text
	checked bool
	selected bool
}

func NewCheckbox(label string, checked bool) *checkbox {
	return &checkbox{
	    label: NewText(label),
	    checked: checked,
	}
}

func (t *checkbox) Text() string {
	return t.label.Text()
}

func (t *checkbox) SetChecked(checked bool) {
	t.checked = checked
}

func (t *checkbox) Render(canvas core.Canvas) {
	w,h := canvas.Size()
	labelCanvas := canvas.Cutout(4, 0, w-4, h)
	t.label.Render(labelCanvas)

	canvas.Set(0,0, '[', nil)
	canvas.Set(2,0, ']', nil)

	if t.checked {
		canvas.Set(1, 0, 'X', nil)
	}
}

func (t *checkbox) Size(parent core.Canvas) (int, int) {
	w, h := t.label.Size(parent)
	return w+4, h
}

func(l *checkbox) ToggleSelect() bool {
	l.selected = !l.selected
	return l.selected
}

func(l *checkbox) HandleKeypress(key *tcell.EventKey) {
	switch key.Key() {
	case tcell.KeyRune:
		switch key.Rune() {
		case ' ':
			l.checked = !l.checked
		}
	}
}
