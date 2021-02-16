package window

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type baseCanvas struct {
	screen tcell.Screen
}

func NewBaseCanvas(screen tcell.Screen) *baseCanvas {
	return &baseCanvas{
		screen: screen,
	}
}

func (c *baseCanvas) Set(x, y int, r rune, s core.Style) {
	c.screen.SetCell(x, y, s.Tcell(), r)
}

func (c *baseCanvas) Size() core.Size {
	w, h := c.screen.Size()
	return core.Size{
		W: w,
		H: h,
	}
}

func (c *baseCanvas) Cutout(x, y int, size core.Size) core.Canvas {
	return NewCutoutCanvas(c, x, y, size)
}

func (c *baseCanvas) Fill(r rune, s core.Style) {
	size := c.Size()
	for x := 0; x < size.W; x++ {
		for y := 0; y < size.H; y++ {
			c.Set(x, y, r, s)
		}
	}
}
