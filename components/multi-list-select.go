package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type listMultiSelect struct {
	core.Sizer
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
	return t
}

func (t *listMultiSelect) WithOptions(options ...string) *listMultiSelect {
	for _, opt := range options {
		t.WithOption(opt)
	}
	return t
}

func (t *listMultiSelect) Render(canvas core.Canvas) {
	var y int
	for _, opt := range t.options {
		actualSize := core.CalculateSize(opt, canvas.Size())
		cutout := canvas.Cutout(0, y, actualSize)
		y += actualSize.H
		opt.Render(cutout)
	}
}

func (t *listMultiSelect) MinimumSize() core.Size {
	var required core.Size
	for _, opt := range t.options {
		optSize := opt.MinimumSize()
		if optSize.W > required.W {
			required.W = optSize.W
		}
		required.H += optSize.H
	}
	return required
}

func (l *listMultiSelect) Deselect() {
	l.selected = false
	if l.selectionIndex < len(l.options) {
		l.options[l.selectionIndex].Deselect()
	}
}

func (l *listMultiSelect) Select() bool {
	if l.selected {
		return false
	}
	l.selected = true
	if l.selectionIndex < len(l.options) {
		l.options[l.selectionIndex].Select()
	}
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
		l.options[l.selectionIndex].Select()
	case tcell.KeyDown:
		l.options[l.selectionIndex].Deselect()
		if l.selectionIndex >= len(l.options)-1 {
			l.selectionIndex = 0
		} else {
			l.selectionIndex++
		}
		l.options[l.selectionIndex].Select()
	case tcell.KeyEnter:
		l.options[l.selectionIndex].checked = !l.options[l.selectionIndex].checked
	case tcell.KeyRune:
		switch key.Rune() {
		case ' ':
			l.options[l.selectionIndex].checked = !l.options[l.selectionIndex].checked
		}
	}
}
