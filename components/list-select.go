package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type listSelect struct {
	core.Sizer
	options        []*checkbox
	selectionIndex int
	selected       bool
}

func NewListSelect(options []string) *listSelect {
	list := &listSelect{}
	return list.WithOptions(options...)
}

// -1 means nothing is selected
func (l *listSelect) GetSelection() (int, string) {
	for i, cb := range l.options {
		if cb.checked {
			return i, cb.Text()
		}
	}
	return -1, ""
}

func (l *listSelect) WithOption(text string) *listSelect {
	cb := NewCheckbox(text, false)
	l.options = append(l.options, cb)
	return l
}

func (l *listSelect) WithOptions(options ...string) *listSelect {
	for _, opt := range options {
		l.WithOption(opt)
	}
	return l
}

func (l *listSelect) Render(canvas core.Canvas) {
	var y int
	for _, opt := range l.options {
		actualSize := core.CalculateSize(opt, canvas.Size())
		cutout := canvas.Cutout(0, y, actualSize)
		y += actualSize.H
		opt.Render(cutout)
	}
}

func (l *listSelect) MinimumSize() core.Size {
	var required core.Size
	for _, opt := range l.options {
		optSize := opt.MinimumSize()
		if optSize.W > required.W {
			required.W = optSize.W
		}
		required.H += optSize.H
	}
	return required
}

func (l *listSelect) Deselect() {
	l.selected = false
	if l.selectionIndex < len(l.options) {
		l.options[l.selectionIndex].Deselect()
	}
}

func (l *listSelect) Select() bool {
	if l.selected {
		return false
	}
	l.selected = true
	if l.selectionIndex < len(l.options) {
		l.options[l.selectionIndex].Select()
	}
	return true
}

func (l *listSelect) HandleKeypress(key *tcell.EventKey) {
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
		l.toggleCurrent()
	case tcell.KeyRune:
		switch key.Rune() {
		case ' ':
			l.toggleCurrent()
		}
	}
}

func (l *listSelect) toggleCurrent() {
	if !l.options[l.selectionIndex].checked {
		for _, opt := range l.options {
			opt.checked = false
		}
	}
	l.options[l.selectionIndex].checked = !l.options[l.selectionIndex].checked
}
