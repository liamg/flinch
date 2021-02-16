package components

import "github.com/liamg/flinch/core"

type testComponent struct {
	size   core.Size
	canvas core.Canvas
}

func (t *testComponent) MinimumSize() core.Size {
	return t.size
}

func newTestComponent(size core.Size, canvas core.Canvas) *testComponent {
	return &testComponent{
		size:   size,
		canvas: canvas,
	}
}

func (t *testComponent) Render(canvas core.Canvas) {
	t.canvas = canvas
}

func (t *testComponent) Size() core.Size {
	return t.size
}

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

func (t *testCanvas) Fill(r rune, st core.Style) {

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

func (t *testCanvas) Cutout(x, y int, size core.Size) core.Canvas {
	return newTestCanvas(x, y, size)
}
