package components

import (
	"github.com/liamg/flinch/core"
)

type columnLayout struct {
	components    []core.Component
	justification core.Justification
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

func (l *columnLayout) SetJustification(justification core.Justification) {
	l.justification = justification
}

func (l *columnLayout) Render(canvas core.Canvas) {

	requiredWidth, _ := l.Size()

	availableWidth, availableHeight := canvas.Size()

	if requiredWidth > availableWidth {
		requiredWidth = availableWidth
	}

	var startX int
	var spacing int

	switch l.justification {
	case core.JustifyLeft:
		startX = 0
		spacing = 0
	case core.JustifyRight:
		startX = availableWidth - requiredWidth
		spacing = 0
	case core.JustifyCenter:
		startX = (availableWidth - requiredWidth) / 2
		spacing = 0
	case core.JustifyFill:
		startX = 0
		spacing = (availableWidth - requiredWidth) / (len(l.components) - 1)
	}

	for _, component := range l.components {
		cWidth, cHeight := component.Size()
		if cHeight > availableHeight {
			cHeight = availableHeight
		}
		if cWidth > availableWidth {
			cWidth = availableWidth
		}
		cutout := canvas.Cutout(startX, 0, cWidth, cHeight)
		component.Render(cutout)
		availableWidth -= cWidth
		startX += spacing
	}
}

func (l *columnLayout) Size() (int, int) {
	var requiredWidth int
	var requiredHeight int
	for _, component := range l.components {
		w, h := component.Size()
		requiredWidth += w
		if h > requiredHeight {
			requiredWidth = h
		}
	}
	return requiredWidth, requiredHeight
}

func NewColumnLayout() *columnLayout {
	return &columnLayout{}
}
