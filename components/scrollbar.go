package components

import "github.com/liamg/flinch/core"

func drawScrollbar(canvas core.Canvas, index int, size int, max int) {

	canvasSize := canvas.Size()
	height := (size * canvasSize.H) / max
	barPos := (index * canvasSize.H) / max

	for y := 0; y < canvasSize.H; y++ {
		if y >= barPos && y < barPos+height {
			canvas.Set(0, y, ' ', core.StyleDefault.ToggleInvert())
		} else {
			canvas.Set(0, y, '│', core.StyleDefault)
		}
	}

}
