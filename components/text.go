package components

import (
	"strings"

	"github.com/liamg/flinch/core"
)

type text struct {
	content       string
	justification core.Justification
}

func NewText(content string) *text {
	return &text{
		content: content,
	}
}

func (t *text) WithJustification(j core.Justification) *text {
	t.justification = j
	return t
}

func (t *text) cleanContent() string {
	return strings.Split(t.content, "\n")[0]
}

func (t *text) Render(canvas core.Canvas) {

	var x int
	w, _ := canvas.Size()

	content := t.cleanContent()

	switch t.justification {
	case core.JustifyLeft:
		x = 0
	case core.JustifyRight:
		x = w - len(content)
	case core.JustifyCenter, core.JustifyFill:
		x = (w - len(content)) / 2
	}

	for offset, r := range content {
		canvas.Set(x+offset, 0, r, nil)
	}
}

func (t *text) Size() (int, int) {
	return len(t.cleanContent()), 1
}
