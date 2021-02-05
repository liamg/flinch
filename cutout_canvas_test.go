package flinch

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liamg/flinch/core"
)

type testCanvas struct {
	cells map[int]map[int]rune
	w     int
	h     int
	x     int
	y     int
}

func newTestCanvas(x, y, w, h int) *testCanvas {
	return &testCanvas{
		x: x,
		y: y,
		w: w,
		h: h,
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

func (t *testCanvas) Size() (w int, h int) {
	return t.w, t.h
}

func (t *testCanvas) Cutout(x, y, w, h int) core.Canvas {
	return newTestCanvas(x, y, w, h)
}

func TestCutoutCanvasSet(t *testing.T) {
	base := newTestCanvas(0, 0, 20, 10)
	cutout := NewCutoutCanvas(base, 10, 5, 10, 5)
	cutout.Set(0, 0, '!', nil)
	assert.Equal(t, '!', base.Get(10, 5))
}

func TestCutoutCanvasSetOutOfBoundsX(t *testing.T) {
	base := newTestCanvas(0, 0, 20, 10)
	cutout := NewCutoutCanvas(base, 10, 5, 10, 5)
	cutout.Set(20, 0, '!', nil)
	assert.NotEqual(t, '!', base.Get(10, 5))
}

func TestCutoutCanvasSetOutOfBoundsY(t *testing.T) {
	base := newTestCanvas(0, 0, 20, 10)
	cutout := NewCutoutCanvas(base, 10, 5, 10, 5)
	cutout.Set(0, 20, '!', nil)
	assert.NotEqual(t, '!', base.Get(10, 5))
}

func TestCutoutCanvasSize(t *testing.T) {
	base := newTestCanvas(0, 0, 20, 10)
	cutout := NewCutoutCanvas(base, 10, 5, 10, 5)
	w, h := cutout.Size()
	assert.Equal(t, 10, w)
	assert.Equal(t, 5, h)
}

func TestCutoutCanvasCutout(t *testing.T) {
	base := newTestCanvas(0, 0, 20, 10)
	cutout := NewCutoutCanvas(base, 10, 5, 10, 5)
	child := cutout.Cutout(5, 2, 3, 2)
	w, h := child.Size()
	assert.Equal(t, 3, w)
	assert.Equal(t, 2, h)
	child.Set(0, 0, '!', nil)
	assert.Equal(t, '!', base.Get(15, 7))
}
