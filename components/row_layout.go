package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type rowLayout struct {
	components    []core.Component
	selector *core.Selector
}


func NewRowLayout() *rowLayout {
	return &rowLayout{
		selector: core.NewSelector(),
	}
}

func (l *rowLayout) Add(component core.Component) {
	for _, comp := range l.components {
		if comp == component {
			return
		}
	}
	l.components = append(l.components, component)
}

func (l *rowLayout) Remove(component core.Component) {
	for i, comp := range l.components {
		if comp == component {
			l.components = append(l.components[:i], l.components[i+1:]...)
			break
		}
	}
}

func (l *rowLayout) WithJustification(_ core.Justification) core.Container {
	return l
}

func (l *rowLayout) Render(canvas core.Canvas) {

	var startY int
	for _, component := range l.components {
		cWidth, cHeight := component.Size(canvas)
		cutout := canvas.Cutout(0, startY, cWidth, cHeight)
		component.Render(cutout)
		startY += cHeight
	}
}

func (l *rowLayout) Size(c core.Canvas) (int, int) {
	var requiredWidth int
	var requiredHeight int
	for _, component := range l.components {
		w, h := component.Size(c)
		requiredHeight += h
		if w > requiredWidth {
			requiredWidth = w
		}
	}
	return requiredWidth, requiredHeight
}


func(l *rowLayout) ToggleSelect() bool {
	return l.selector.ToggleSelect(l.components)
}

func(l *rowLayout) HandleKeypress(key *tcell.EventKey) {
	sel := l.selector.GetSelected()
	if sel != nil {
		sel.HandleKeypress(key)
	}
}