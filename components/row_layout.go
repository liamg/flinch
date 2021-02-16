package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type rowLayout struct {
	components    []core.Component
	justification core.Justification
	selector      *core.Selector
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

func (l *rowLayout) WithJustification(justification core.Justification) core.Container {
	l.justification = justification
	return l
}

func (l *rowLayout) Render(canvas core.Canvas) {

	_, requiredHeight := l.Size(canvas)

	availableWidth, availableHeight := canvas.Size()

	if requiredHeight > availableWidth {
		requiredHeight = availableHeight
	}

	var startY int
	var spacing int

	switch l.justification {
	case core.JustifyLeft:
		startY = 0
		spacing = 0
	case core.JustifyRight:
		startY = availableHeight - requiredHeight
		spacing = 0
	case core.JustifyCenter:
		startY = (availableHeight - requiredHeight) / 2
		spacing = 0
	case core.JustifyFill:
		startY = 0
		spacing = (availableHeight - requiredHeight) / len(l.components)
	}

	for _, component := range l.components {
		_, cHeight := component.Size(canvas)
		cHeight = cHeight + spacing
		if cHeight > availableHeight {
			cHeight = availableHeight
		}
		cutout := canvas.Cutout(0, startY, availableWidth, cHeight)
		component.Render(cutout)
		availableHeight -= cHeight
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

func (l *rowLayout) Deselect() {
	l.selector.Deselect()
}

func (l *rowLayout) Select(loop bool) bool {
	return l.selector.Select(l.components, loop)
}

func (l *rowLayout) HandleKeypress(key *tcell.EventKey) {
	sel := l.selector.GetSelected()
	if sel != nil {
		sel.HandleKeypress(key)
	}
}
