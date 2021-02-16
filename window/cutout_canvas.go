package window

import (
	"github.com/liamg/flinch/core"
)

type cutoutCanvas struct {
	parent core.Canvas
	x      int
	y      int
	size core.Size
}

func NewCutoutCanvas(parent core.Canvas, x, y int, size core.Size) *cutoutCanvas {
	return &cutoutCanvas{
		parent: parent,
		x:      x,
		y:      y,
		size: size,
	}
}

func (c *cutoutCanvas) Set(x, y int, r rune, s core.Style) {

	if x >= c.size.W {
		return
	}
	if y >= c.size.H {
		return
	}
	tX := c.x + x
	tY := c.y + y

	c.parent.Set(tX, tY, r, s)
}

func (c *cutoutCanvas) Size() core.Size {
	return c.size
}

func (c *cutoutCanvas) Cutout(x, y int, size core.Size) core.Canvas {
	return NewCutoutCanvas(c, x, y, size)
}

func (c *cutoutCanvas) Fill(r rune, s core.Style) {
	for x := 0; x < c.size.W; x++ {
		for y := 0; y < c.size.H; y++ {
			c.Set(x, y, r, s)
		}
	}
}
