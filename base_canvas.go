package flinch

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
	// TODO convert style
	c.screen.SetCell(x, y, tcell.StyleDefault, r)
}

func (c *baseCanvas) Size() (w int, h int) {
	return c.screen.Size()
}

func (c *baseCanvas) Cutout(x, y, w, h int) core.Canvas {
	return NewCutoutCanvas(c, x, y, w, h)
}
