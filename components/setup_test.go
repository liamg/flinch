package components

import "github.com/liamg/flinch/core"

type testComponent struct {
	w      int
	h      int
	canvas core.Canvas
}

func newTestComponent(w, h int, canvas core.Canvas) *testComponent {
	return &testComponent{
		w:      w,
		h:      h,
		canvas: canvas,
	}
}

func (t *testComponent) SetSize(w, h int) {
	t.w = w
	t.h = h
}

func (t *testComponent) Render(canvas core.Canvas) {
	t.canvas = canvas
}

func (t *testComponent) Size() (int, int) {
	return t.w, t.h
}

type testCanvas struct {
	cells map[int]map[int]rune
	w     int
	h     int
}

func newTestCanvas(w, h int) *testCanvas {
	return &testCanvas{
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
	return nil
}
