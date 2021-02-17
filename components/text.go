package components

import (
	"strings"

	"github.com/liamg/flinch/core"
)

type text struct {
	core.Sizer
	content   string
	alignment core.Alignment
	style     core.Style
	padSize   int
}

func NewText(content string) *text {
	return &text{
		content: content,
		style:   core.StyleDefault,
	}
}

func (t *text) PadText(size int) {
	t.padSize = size
}

func (t *text) SetAlignment(j core.Alignment) *text {
	t.alignment = j
	return t
}

func (t *text) Text() string {
	return t.content
}

func (t *text) SetStyle(s core.Style) {
	t.style = s
}

func (t *text) cleanContent() string {
	return strings.Split(t.content, "\n")[0]
}

func (t *text) Render(canvas core.Canvas) {

	var x int
	size := canvas.Size()

	content := t.cleanContent()

	switch t.alignment {
	case core.AlignLeft:
		x = 0
	case core.AlignRight:
		x = size.W - len(content)
	case core.AlignCenter:
		x = (size.W - len(content)) / 2
	}

	for offset, r := range content {
		canvas.Set(x+offset+t.padSize, 0, r, t.style)
	}
}

func (t *text) MinimumSize() core.Size {
	return core.Size{
		W: len(t.cleanContent()),
		H: 1,
	}
}
