package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type limiter struct {
	inner       core.Component
	pcLimitW    int
	pcLimitH    int
	fixedLimitW int
	fixedLimitH int
	minW        int
	minH        int
}

func NewLimiter(inner core.Component) *limiter {
	return &limiter{
		inner: inner,
	}
}

func (t *limiter) WithMinumumSize(w, h int) *limiter {
	t.minW = w
	t.minH = h
	return t
}

func (t *limiter) WithPercentageLimitOnWidth(pc int) *limiter {
	t.pcLimitW = pc
	t.fixedLimitW = 0
	return t
}

func (t *limiter) WithPercentageLimitOnHeight(pc int) *limiter {
	t.pcLimitH = pc
	t.fixedLimitH = 0
	return t
}

func (t *limiter) WithFixedLimitOnWidth(w int) *limiter {
	t.pcLimitW = 0
	t.fixedLimitW = w
	return t
}

func (t *limiter) WithFixedLimitOnHeight(h int) *limiter {
	t.pcLimitH = 0
	t.fixedLimitH = h
	return t
}

func (t *limiter) Render(canvas core.Canvas) {
	t.inner.Render(canvas)
}

func (t *limiter) Size(parent core.Canvas) (int, int) {

	pw, ph := t.inner.Size(parent)

	if t.fixedLimitW > 0 {
		if pw > t.fixedLimitW {
			pw = t.fixedLimitW
		}
	} else if t.pcLimitW > 0 {
		pw = (pw * t.pcLimitW) / 100
	}

	if t.fixedLimitH > 0 {
		if ph > t.fixedLimitH {
			ph = t.fixedLimitH
		}
	} else if t.pcLimitH > 0 {
		ph = (ph * t.pcLimitH) / 100
	}

	if ph < t.minH {
		ph = t.minH
	}
	if pw < t.minW {
		pw = t.minW
	}

	return pw, ph
}

func (l *limiter) ToggleSelect(loop bool) bool {
	if sel, ok := l.inner.(core.Selectable); ok {
		return sel.ToggleSelect(loop)
	}
	return false
}

func (l *limiter) HandleKeypress(key *tcell.EventKey) {
	if sel, ok := l.inner.(core.Selectable); ok {
		sel.HandleKeypress(key)
	}
}
