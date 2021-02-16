package core

type Canvas interface {
	Fill(r rune, s Style)
	Set(x, y int, r rune, s Style)
	Size() Size
	Cutout(x, y int, s Size) Canvas
}
