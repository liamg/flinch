package components

import (
	"github.com/liamg/flinch/core"
)

type listSelect struct {
	options []*checkbox
}

func NewListSelect(label string, checked bool) *listSelect {
	return &listSelect{
	}
}

func(t *listSelect) GetSelection() ([]int, []string){
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

func(t *listSelect) WithOption(text string) *listSelect {
	cb := NewCheckbox(text, false)
	t.options = append(t.options, cb)
	return t
}



func (t *listSelect) Render(canvas core.Canvas) {
}

func (t *listSelect) Size(parent core.Canvas) (int, int) {
	return parent.Size()
}
