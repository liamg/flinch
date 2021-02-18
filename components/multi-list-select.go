package components

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type listMultiSelect struct {
	core.Sizer
	options        []*checkbox
	selectionIndex int
	selected       bool
	keyHandlers    []func(key *tcell.EventKey) bool
	viewIndex      int
}

func NewMultiListSelect(options []string) *listMultiSelect {
	list := &listMultiSelect{}
	return list.WithOptions(options...)
}

func (m *listMultiSelect) GetSelection() ([]int, []string) {
	var indexes []int
	var strings []string
	for i, cb := range m.options {
		if cb.checked {
			indexes = append(indexes, i)
			strings = append(strings, cb.Text())
		}
	}
	return indexes, strings
}

func (m *listMultiSelect) WithOption(text string) *listMultiSelect {
	cb := NewCheckbox(text, false)
	m.options = append(m.options, cb)
	return m
}

func (m *listMultiSelect) WithOptions(options ...string) *listMultiSelect {
	for _, opt := range options {
		m.WithOption(opt)
	}
	return m
}

func (m *listMultiSelect) Render(canvas core.Canvas) {
	var y int
	size := canvas.Size()
	visibleCount := size.H
	size.H = 1

	if m.selectionIndex < m.viewIndex {
		m.viewIndex = m.selectionIndex
	} else if m.selectionIndex >= m.viewIndex+visibleCount {
		m.viewIndex = (m.selectionIndex - visibleCount) + 1
	}

	if len(m.options) > visibleCount {
		size.W--
		scrollCanvas := canvas.Cutout(size.W, 0, core.Size{W: 1, H: visibleCount})
		drawScrollbar(scrollCanvas, m.viewIndex, visibleCount, len(m.options))
	}

	for index := m.viewIndex; index < m.viewIndex+visibleCount && index < len(m.options); index++ {
		opt := m.options[index]
		actualSize := core.CalculateSize(opt, size)
		cutout := canvas.Cutout(0, y, actualSize)
		y += actualSize.H
		opt.Render(cutout)
	}
}

func (l *listMultiSelect) MinimumSize() core.Size {
	maxVisible := 10
	if len(l.options) < maxVisible {
		maxVisible = len(l.options)
	}
	return core.Size{
		W: 10,
		H: maxVisible,
	}
}

func (m *listMultiSelect) Deselect() {
	m.selected = false
	if m.selectionIndex < len(m.options) {
		m.options[m.selectionIndex].Deselect()
	}
}

func (m *listMultiSelect) Select() bool {
	if m.selected {
		return false
	}
	m.selected = true
	if m.selectionIndex < len(m.options) {
		m.options[m.selectionIndex].Select()
	}
	return true
}

func (m *listMultiSelect) OnKeypress(handler func(key *tcell.EventKey) bool) {
	m.keyHandlers = append(m.keyHandlers, handler)
}

func (m *listMultiSelect) HandleKeypress(key *tcell.EventKey) {

	for _, handler := range m.keyHandlers {
		if handler(key) {
			return
		}
	}

	switch key.Key() {
	case tcell.KeyUp:
		m.options[m.selectionIndex].Deselect()
		if m.selectionIndex > 0 {
			m.selectionIndex--
		}
		m.options[m.selectionIndex].Select()
	case tcell.KeyDown:
		m.options[m.selectionIndex].Deselect()
		if m.selectionIndex < len(m.options)-1 {
			m.selectionIndex++
		}
		m.options[m.selectionIndex].Select()
	case tcell.KeyEnter:
		m.options[m.selectionIndex].checked = !m.options[m.selectionIndex].checked
	case tcell.KeyRune:
		switch key.Rune() {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			index, _ := strconv.Atoi(string([]rune{key.Rune()}))
			index = index - 1
			if index < len(m.options) {
				m.selectionIndex = index
			}
		case ' ':
			m.options[m.selectionIndex].checked = !m.options[m.selectionIndex].checked
		}
	}
}
