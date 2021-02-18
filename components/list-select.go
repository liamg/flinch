package components

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type listSelect struct {
	core.Sizer
	options        []string
	selectionIndex int
	viewIndex      int
	selected       bool
	keyHandlers    []func(key *tcell.EventKey) bool
}

func NewListSelect(options []string) *listSelect {
	list := &listSelect{}
	return list.WithOptions(options...)
}

// -1 means nothing is selected
func (l *listSelect) GetSelection() (int, string) {
	if l.selectionIndex > len(l.options)-1 {
		return -1, ""
	}
	return l.selectionIndex, l.options[l.selectionIndex]
}

func (l *listSelect) WithOption(text string) *listSelect {
	l.options = append(l.options, text)
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
	size := canvas.Size()
	visibleCount := size.H
	size.H = 1

	if l.selectionIndex < l.viewIndex {
		l.viewIndex = l.selectionIndex
	} else if l.selectionIndex >= l.viewIndex+visibleCount {
		l.viewIndex = (l.selectionIndex - visibleCount) + 1
	}

	if len(l.options) > visibleCount {
		size.W--
		scrollCanvas := canvas.Cutout(size.W, 0, core.Size{W: 1, H: visibleCount})
		drawScrollbar(scrollCanvas, l.viewIndex, visibleCount, len(l.options))
	}

	for index := l.viewIndex; index < l.viewIndex+visibleCount && index < len(l.options); index++ {
		opt := l.options[index]
		cutout := canvas.Cutout(0, y, size)
		st := core.StyleDefault
		if index == l.selectionIndex {
			st = core.StyleSelected
		}
		y++
		cutout.Fill(' ', st)
		if index == l.selectionIndex {
			cutout.Set(1, 0, '✔', st)
		}
		for i, char := range []rune(opt) {
			cutout.Set(i+3, 0, char, st)
		}
	}
}

func (l *listSelect) MinimumSize() core.Size {
	maxVisible := 10
	if len(l.options) < maxVisible {
		maxVisible = len(l.options)
	}
	return core.Size{
		W: 10,
		H: maxVisible,
	}
}

func (l *listSelect) Deselect() {
	l.selected = false
}

func (l *listSelect) Select() bool {
	if l.selected {
		return false
	}
	l.selected = true
	return true
}

func (n *listSelect) OnKeypress(handler func(key *tcell.EventKey) bool) {
	n.keyHandlers = append(n.keyHandlers, handler)
}

func (l *listSelect) HandleKeypress(key *tcell.EventKey) {

	for _, handler := range l.keyHandlers {
		if handler(key) {
			return
		}
	}

	switch key.Key() {
	case tcell.KeyRune:
		switch key.Rune() {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			index, _ := strconv.Atoi(string([]rune{key.Rune()}))
			index = index - 1
			if index < len(m.options) {
				m.selectionIndex = index
			}
		}
	case tcell.KeyUp:
		if l.selectionIndex > 0 {
			l.selectionIndex--
		}
	case tcell.KeyDown:
		if l.selectionIndex < len(l.options)-1 {
			l.selectionIndex++
		}
	}
}
