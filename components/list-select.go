package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type listMultiSelect struct {
	options        []*checkbox
	selectionIndex int
	selected       bool
}

func NewMultiListSelect(options []string) *listMultiSelect {
	list := &listMultiSelect{}
	return list.WithOptions(options...)
}

func (t *listMultiSelect) GetSelection() ([]int, []string) {
	var indexes []int
	var strings []string
	for i, cb := range t.options {
		if cb.checked {
			indexes = append(indexes, i)
			strings = append(strings, cb.Text())
		}
	}
	return indexes, strings
}

func (t *listMultiSelect) WithOption(text string) *listMultiSelect {
	cb := NewCheckbox(text, false)
	t.options = append(t.options, cb)
	if len(t.options) == 1 {
		cb.Select(false)
	}
	return t
}

func (t *listMultiSelect) WithOptions(options ...string) *listMultiSelect {
	for _, opt := range options {
		t.WithOption(opt)
	}
	return t
}

func (t *listMultiSelect) Render(canvas core.Canvas) {
	w, h := t.Size(canvas)
	for x := 0; x < w; x++ {
		canvas.Set(x, 0, getBorderRune(x, 0, w, h), core.StyleDefault)
		canvas.Set(x, h-1, getBorderRune(x, h-1, w, h), core.StyleDefault)
	}
	for y := 0; y < h; y++ {
		canvas.Set(0, y, getBorderRune(0, y, w, h), core.StyleDefault)
		canvas.Set(w-1, y, getBorderRune(w-1, y, w, h), core.StyleDefault)
	}

	for y, opt := range t.options {
		w, h := opt.Size(canvas)
		cutout := canvas.Cutout(1, y+1, w-2, h)
		opt.Render(cutout)
	}
}

func (t *listMultiSelect) Size(parent core.Canvas) (int, int) {
	w, _ := parent.Size()
	return w, len(t.options) + 2
}

func (l *listMultiSelect) Deselect() {
	l.selected = false
}

func (l *listMultiSelect) Select(loop bool) bool {
	if l.selected {
		return false
	}
	l.selected = true
	return true
}

func (l *listMultiSelect) HandleKeypress(key *tcell.EventKey) {
	switch key.Key() {
	case tcell.KeyUp:
		l.options[l.selectionIndex].Deselect()
		if l.selectionIndex <= 0 {
			l.selectionIndex = len(l.options) - 1
		} else {
			l.selectionIndex--
		}
		l.options[l.selectionIndex].Select(false)
	case tcell.KeyDown:
		l.options[l.selectionIndex].Deselect()
		if l.selectionIndex >= len(l.options)-1 {
			l.selectionIndex = 0
		} else {
			l.selectionIndex++
		}
		l.options[l.selectionIndex].Select(false)
	case tcell.KeyEnter:
		l.options[l.selectionIndex].checked = !l.options[l.selectionIndex].checked
	case tcell.KeyRune:
		switch key.Rune() {
		case ' ':
			l.options[l.selectionIndex].checked = !l.options[l.selectionIndex].checked
		}
	}
}
