package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type columnLayout struct {
	core.Sizer
	components        []core.Component
	justification     core.Alignment
	selectedComponent core.Selectable
	selector          *core.Selector
}

func NewColumnLayout() *columnLayout {
	layout := &columnLayout{
		selector: core.NewSelector(),
	}
	layout.SetSizeStrategy(core.SizeStrategyMaximum())
	return layout
}

func (l *columnLayout) Add(component core.Component) {
	for _, comp := range l.components {
		if comp == component {
			return
		}
	}
	l.components = append(l.components, component)
}

func (l *columnLayout) Remove(component core.Component) {
	for i, comp := range l.components {
		if comp == component {
			l.components = append(l.components[:i], l.components[i+1:]...)
			break
		}
	}
}

func (l *columnLayout) SetAlignment(justification core.Alignment) {
	l.justification = justification
}

func (l *columnLayout) Render(canvas core.Canvas) {

	availableSize := canvas.Size()
	requiredSize := l.MinimumSize()
	extraSpace := availableSize.Minus(requiredSize)

	var usedWidth int
	var actualWidth int
	for i, component := range l.components {
		spacing := (extraSpace.W * requiredSize.W) / availableSize.W
		min := component.MinimumSize()
		if i == len(l.components)-1 {
			spacing = (availableSize.W - usedWidth) - min.W
		}
		compAvail := core.Size{W: min.W + spacing, H: availableSize.H}
		compSize := core.CalculateSize(component, compAvail)
		actualWidth += compSize.W
	}

	for i, component := range l.components {
		spacing := (extraSpace.W * requiredSize.W) / availableSize.W
		min := component.MinimumSize()
		if i == len(l.components)-1 {
			spacing = (availableSize.W - usedWidth) - min.W
		}
		compAvail := core.Size{W: min.W + spacing, H: availableSize.H}
		compSize := core.CalculateSize(component, compAvail)

		var cutout core.Canvas
		switch l.justification {
		case core.AlignLeft:
			cutout = canvas.Cutout(usedWidth, 0, compSize)
		case core.AlignRight:
			cutout = canvas.Cutout((availableSize.W-actualWidth)-usedWidth, 0, compSize)
		case core.AlignCenter:
			cutout = canvas.Cutout(usedWidth+((availableSize.W-actualWidth)/2), 0, compSize)
		}

		usedWidth += compSize.W
		component.Render(cutout)
	}
}

func (t *columnLayout) MinimumSize() core.Size {
	var required core.Size
	for _, comp := range t.components {
		min := comp.MinimumSize()
		required.W += min.W
		if min.H > required.H {
			required.H = min.H
		}
	}
	return required
}

func (l *columnLayout) Deselect() {
	l.selector.Deselect()
}

func (l *columnLayout) Select() bool {
	return l.selector.Select(l.components)
}

func (l *columnLayout) HandleKeypress(key *tcell.EventKey) {
	sel := l.selector.GetSelected()
	if sel != nil {
		sel.HandleKeypress(key)
	}
}
