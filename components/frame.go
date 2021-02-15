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

	for x := 0 ; x < w; x++ {
		canvas.Set(x, 0, t.getRune(x,0,w,h), nil)
		canvas.Set(x, h-1, t.getRune(x,h-1,w,h), nil)
	}
	for y := 0 ; y < h; y++ {
		canvas.Set(0, y, t.getRune(0,y,w,h), nil)
		canvas.Set(w-1, y, t.getRune(w-1,y,w,h), nil)
	}

	innerCanvas := canvas.Cutout(1, 1, w-2, h-2)
	t.inner.Render(innerCanvas)
}

func(t *frame) getRune(x, y, w, h int) rune {
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

func(l *frame) ToggleSelect() bool {
	if sel, ok := l.inner.(core.Selectable); ok {
		return sel.ToggleSelect()
	}
	return false
}

func(l *frame) HandleKeypress(key *tcell.EventKey) {
	if sel, ok := l.inner.(core.Selectable); ok {
		sel.HandleKeypress(key)
	}
}
