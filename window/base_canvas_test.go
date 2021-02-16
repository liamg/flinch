package window

import (
	"testing"

	"github.com/liamg/flinch/core"

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
	canvas.Set(5, 5, 'x', core.StyleDefault)

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
	cSize := canvas.Size()

	assert.Equal(t, eW, cSize.W)
	assert.Equal(t, eH, cSize.H)
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
	child := canvas.Cutout(0, 0, core.Size{W: 5, H: 10})

	cSize := child.Size()
	assert.Equal(t, 5, cSize.W)
	assert.Equal(t, 10, cSize.H)
}
