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

	size := canvas.Size()

	for x := 0; x < size.W; x++ {
		canvas.Set(x, 0, getBorderRune(x, 0, size.W, size.H), core.StyleDefault)
		canvas.Set(x, size.H-1, getBorderRune(x, size.H-1, size.W, size.H), core.StyleDefault)
	}
	for y := 0; y < size.H; y++ {
		canvas.Set(0, y, getBorderRune(0, y, size.W, size.H), core.StyleDefault)
		canvas.Set(size.W-1, y, getBorderRune(size.W-1, y, size.W, size.H), core.StyleDefault)
	}

	innerCanvas := canvas.Cutout(1, 1, size.Minus(core.Size{W: 2, H: 2}))
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

func (t *frame) MinimumSize() core.Size {
	return t.inner.MinimumSize().Add(core.Size{W: 2, H: 2})
}

func (t *frame) Deselect() {
	if sel, ok := t.inner.(core.Selectable); ok {
		sel.Deselect()
	}
}

func (t *frame) Select() bool {
	if sel, ok := t.inner.(core.Selectable); ok {
		return sel.Select()
	}
	return false
}

func (l *frame) HandleKeypress(key *tcell.EventKey) {
	if sel, ok := l.inner.(core.Selectable); ok {
		sel.HandleKeypress(key)
	}
}

func (s *frame) SetSizeStrategy(strategy core.SizeStrategy) {
	if sizer, ok := s.inner.(core.StrategicSizer); ok {
		sizer.SetSizeStrategy(strategy)
	}
}

func (s frame) GetSizeStrategy() core.SizeStrategy {
	if sizer, ok := s.inner.(core.StrategicSizer); ok {
		return sizer.GetSizeStrategy()
	}
	return core.SizeStrategyMaximum()
}
