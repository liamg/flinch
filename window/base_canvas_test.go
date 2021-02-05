package window

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
)

func TestBaseCanvasSet(t *testing.T) {
	screen, err := tcell.NewScreen()
	if err != nil {
		t.Fatal(err)
	}
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	defer screen.Fini()

	canvas := NewBaseCanvas(screen)
	canvas.Set(5, 5, 'x', nil)

	r, _, _, _ := screen.GetContent(5, 5)

	assert.Equal(t, 'x', r)

}

func TestBaseCanvasSize(t *testing.T) {
	screen, err := tcell.NewScreen()
	if err != nil {
		t.Fatal(err)
	}
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	defer screen.Fini()

	canvas := NewBaseCanvas(screen)

	eW, eH := screen.Size()
	aW, aH := canvas.Size()

	assert.Equal(t, eW, aW)
	assert.Equal(t, eH, aH)
}

func TestBaseCanvasCutout(t *testing.T) {
	screen, err := tcell.NewScreen()
	if err != nil {
		t.Fatal(err)
	}
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	defer screen.Fini()

	canvas := NewBaseCanvas(screen)
	child := canvas.Cutout(0, 0, 5, 10)

	w, h := child.Size()
	assert.Equal(t, 5, w)
	assert.Equal(t, 10, h)
}
