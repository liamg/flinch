package flinch

import (
	"github.com/liamg/flinch/core"
)

type cutoutCanvas struct {
	parent core.Canvas
	x      int
	y      int
	w      int
	h      int
}

func NewCutoutCanvas(parent core.Canvas, x, y, w, h int) *cutoutCanvas {
	return &cutoutCanvas{
		parent: parent,
		x:      x,
		y:      y,
		w:      w,
		h:      h,
	}
}

func (c *cutoutCanvas) Set(x, y int, r rune, s core.Style) {

	if x >= c.w {
		return
	}
	if y >= c.h {
		return
	}
	tX := c.x + x
	tY := c.y + y

	c.parent.Set(tX, tY, r, s)
}

func (c *cutoutCanvas) Size() (w int, h int) {
	return c.w, c.h
}

func (c *cutoutCanvas) Cutout(x, y, w, h int) core.Canvas {
	return NewCutoutCanvas(c, x, y, w, h)
}
