package core

type Canvas interface {
	Set(x, y int, r rune, s Style)
	Size() (w int, h int)
	Cutout(x, y, w, h int) Canvas
}
