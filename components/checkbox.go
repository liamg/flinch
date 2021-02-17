package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type checkbox struct {
	core.Sizer
	label    string
	checked  bool
	selected bool
}

func NewCheckbox(label string, checked bool) *checkbox {
	cb := &checkbox{
		label:   label,
		checked: checked,
	}
	cb.SetSizeStrategy(core.SizeStrategyMaximumWidth())
	return cb
}

func (t *checkbox) Text() string {
	return t.label
}

func (t *checkbox) SetChecked(checked bool) {
	t.checked = checked
}

func (t *checkbox) Render(canvas core.Canvas) {

	st := core.StyleDefault
	faint := core.StyleFaint
	if t.selected {
		st = core.StyleSelected
		faint = st
	}

	canvas.Fill(' ', st)

	canvas.Set(0, 0, '[', faint)
	canvas.Set(2, 0, ']', faint)

	if t.checked {
		canvas.Set(1, 0, '✔', st)
	} else {
		canvas.Set(1, 0, ' ', st)
	}

	for i := 0; i < len(t.label); i++ {
		canvas.Set(4+i, 0, rune(t.label[i]), st)
	}
}

func (t *checkbox) MinimumSize() core.Size {
	rw, rh := len(t.label)+4, 1
	return core.Size{W: rw, H: rh}
}

func (l *checkbox) Deselect() {
	l.selected = false
}

func (l *checkbox) Select() bool {
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
