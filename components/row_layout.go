package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type rowLayout struct {
	core.Sizer
	components    []core.Component
	justification core.Alignment
	selector      *core.Selector
}

func NewRowLayout() *rowLayout {
	layout := &rowLayout{
		selector: core.NewSelector(),
	}
	layout.SetSizeStrategy(core.SizeStrategyMaximum())
	return layout
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

func (l *rowLayout) SetAlignment(justification core.Alignment) {
	l.justification = justification
}

func (l *rowLayout) Render(canvas core.Canvas) {

	availableSize := canvas.Size()
	requiredSize := l.MinimumSize()
	extraSpace := availableSize.Minus(requiredSize)

	var usedHeight int

	for _, component := range l.components {
		spacing := (extraSpace.H) / len(l.components)
		min := component.MinimumSize()
		compAvail := core.Size{W: availableSize.W, H: min.H + spacing}
		compSize := core.CalculateSize(component, compAvail)

		var cutout core.Canvas
		switch l.justification {
		case core.AlignTop:
			cutout = canvas.Cutout(0, usedHeight, compSize)
		case core.AlignBottom:
			cutout = canvas.Cutout(0, (availableSize.H-requiredSize.H)-usedHeight, compSize)
		case core.AlignCenter:
			x := (availableSize.W - compSize.W) / 2
			cutout = canvas.Cutout(x, usedHeight+((availableSize.H-requiredSize.H)/2), compSize)
		}

		usedHeight += compSize.H
		component.Render(cutout)
	}
}

func (t *rowLayout) MinimumSize() core.Size {
	var required core.Size
	for _, comp := range t.components {
		min := comp.MinimumSize()
		required.H += min.H
		if min.W > required.W {
			required.W = min.W
		}
	}
	return required
}

func (l *rowLayout) Deselect() {
	l.selector.Deselect()
}

func (l *rowLayout) Select() bool {
	return l.selector.Select(l.components)
}

func (l *rowLayout) HandleKeypress(key *tcell.EventKey) {
	sel := l.selector.GetSelected()
	if sel != nil {
		sel.HandleKeypress(key)
	}
}
