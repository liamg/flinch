package window

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liamg/flinch/core"
)

type testCanvas struct {
	cells map[int]map[int]rune
	x     int
	y     int
	size  core.Size
}

func newTestCanvas(x, y int, size core.Size) *testCanvas {
	return &testCanvas{
		x:    x,
		y:    y,
		size: size,
	}
}

func (t *testCanvas) Set(x, y int, r rune, s core.Style) {
	if t.cells == nil {
		t.cells = make(map[int]map[int]rune)
	}
	if _, ok := t.cells[x]; !ok {
		t.cells[x] = make(map[int]rune)
	}
	t.cells[x][y] = r
}

func (t *testCanvas) Get(x, y int) rune {
	col, ok := t.cells[x]
	if !ok {
		return 0
	}
	cell, ok := col[y]
	if !ok {
		return 0
	}
	return cell
}

func (t *testCanvas) Size() core.Size {
	return t.size
}

func (t *testCanvas) Fill(_ rune, _ core.Style) {

}

func (t *testCanvas) Cutout(x, y int, size core.Size) core.Canvas {
	return newTestCanvas(x, y, size)
}

func TestCutoutCanvasSet(t *testing.T) {
	base := newTestCanvas(0, 0, core.Size{W: 20, H: 10})
	cutout := NewCutoutCanvas(base, 10, 5, core.Size{W: 10, H: 5})
	cutout.Set(0, 0, '!', core.StyleDefault)
	assert.Equal(t, '!', base.Get(10, 5))
}

func TestCutoutCanvasSetOutOfBoundsX(t *testing.T) {
	base := newTestCanvas(0, 0, core.Size{W: 20, H: 10})
	cutout := NewCutoutCanvas(base, 10, 5, core.Size{W: 10, H: 5})
	cutout.Set(20, 0, '!', core.StyleDefault)
	assert.NotEqual(t, '!', base.Get(10, 5))
}

func TestCutoutCanvasSetOutOfBoundsY(t *testing.T) {
	base := newTestCanvas(0, 0, core.Size{W: 20, H: 10})
	cutout := NewCutoutCanvas(base, 10, 5, core.Size{W: 10, H: 5})
	cutout.Set(0, 20, '!', core.StyleDefault)
	assert.NotEqual(t, '!', base.Get(10, 5))
}

func TestCutoutCanvasSize(t *testing.T) {
	base := newTestCanvas(0, 0, core.Size{W: 20, H: 10})
	cutout := NewCutoutCanvas(base, 10, 5, core.Size{W: 10, H: 5})
	size := cutout.Size()
	assert.Equal(t, 10, size.W)
	assert.Equal(t, 5, size.H)
}

func TestCutoutCanvasCutout(t *testing.T) {
	base := newTestCanvas(0, 0, core.Size{W: 20, H: 10})
	cutout := NewCutoutCanvas(base, 10, 5, core.Size{W: 10, H: 5})
	child := cutout.Cutout(5, 2, core.Size{W: 3, H: 2})
	size := child.Size()
	assert.Equal(t, 3, size.W)
	assert.Equal(t, 2, size.H)
	child.Set(0, 0, '!', core.StyleDefault)
	assert.Equal(t, '!', base.Get(15, 7))
}
