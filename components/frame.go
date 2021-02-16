package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type frame struct {
	inner core.Component
}

func NewFrame(inner core.Component) *frame {
	return &frame{
		inner: inner,
	}
}

func (t *frame) Render(canvas core.Canvas) {
	w, h := canvas.Size()

	for x := 0; x < w; x++ {
		canvas.Set(x, 0, getBorderRune(x, 0, w, h), core.StyleDefault)
		canvas.Set(x, h-1, getBorderRune(x, h-1, w, h), core.StyleDefault)
	}
	for y := 0; y < h; y++ {
		canvas.Set(0, y, getBorderRune(0, y, w, h), core.StyleDefault)
		canvas.Set(w-1, y, getBorderRune(w-1, y, w, h), core.StyleDefault)
	}

	innerCanvas := canvas.Cutout(1, 1, w-2, h-2)
	t.inner.Render(innerCanvas)
}

func getBorderRune(x, y, w, h int) rune {
	var r rune
	switch true {
	case x == 0 && y == 0:
		r = '┌'
	case x == 0 && y == h-1:
		r = '└'
	case (x == 0 || x == w-1) && y > 0 && y < h-1:
		r = '│'
	case x == w-1 && y == 0:
		r = '┐'
	case x == w-1 && y == h-1:
		r = '┘'
	case (y == 0 || y == h-1) && x > 0 && x < w-1:
		r = '─'
	default:
		r = ' '
	}
	return r
}

func (t *frame) Size(parent core.Canvas) (int, int) {
	w, h := t.inner.Size(parent)
	return w + 2, h + 2
}

func (t *frame) Deselect() {
	if sel, ok := t.inner.(core.Selectable); ok {
		sel.Deselect()
	}
}

func (t *frame) Select(loop bool) bool {
	if sel, ok := t.inner.(core.Selectable); ok {
		return sel.Select(loop)
	}
	return false
}

func (l *frame) HandleKeypress(key *tcell.EventKey) {
	if sel, ok := l.inner.(core.Selectable); ok {
		sel.HandleKeypress(key)
	}
}
