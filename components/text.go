package components

import "github.com/liamg/flinch/core"

type text struct {
	content       string
	justification core.Justification
}

func NewText(content string) *text {
	return &text{
		content: content,
	}
}

func (t *text) SetJustification(j core.Justification) {
	t.justification = j
}

func (t *text) Render(canvas core.Canvas) {

	var x int
	w, _ := canvas.Size()

	switch t.justification {
	case core.JustifyLeft, core.JustifyFill:
		x = 0
	case core.JustifyRight:
		x = w - len(t.content)
	case core.JustifyCenter:
		x = (w - len(t.content)) / 2
	}

	for offset, r := range t.content {
		canvas.Set(x+offset, 0, r, core.StyleDefault)
	}
}

func (t *text) Size() (int, int) {
	return len(t.content), 1
}
