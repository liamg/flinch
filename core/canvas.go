package core

type Canvas interface {
	Fill(r rune, s Style)
	Set(x, y int, r rune, s Style)
	Size() (w int, h int)
	Cutout(x, y, w, h int) Canvas
}
