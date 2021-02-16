package core

import "github.com/gdamore/tcell/v2"

type Style struct {
	bg Colour
	fg Colour
}

var StyleDefault = Style{
	bg: Colour([3]int32{0x0, 0x0, 0x44}),
	fg: Colour([3]int32{0xff, 0xff, 0xff}),
}

var StyleFaint = Style{
	bg: StyleDefault.GetBackground(),
	fg: Colour([3]int32{0x80, 0x80, 0x80}),
}

var StyleSelected = Style{
	bg: Colour([3]int32{0xff, 0xff, 0xff}),
	fg: Colour([3]int32{0x00, 0x00, 0x88}),
}

var StyleButton = Style{
	bg: Colour([3]int32{0x80, 0x80, 0x80}),
	fg: StyleDefault.GetBackground(),
}

var StyleButtonSelected = StyleDefault.Invert()

func (s *Style) Invert() Style {
	return Style{
		fg: s.GetBackground(),
		bg: s.GetForeground(),
	}
}

func (s *Style) GetForeground() Colour {
	return s.fg
}

func (s *Style) GetBackground() Colour {
	return s.bg
}

func (s *Style) SetForeground(colour Colour) {
	s.fg = colour
}

func (s *Style) SetBackground(colour Colour) {
	s.bg = colour
}

func (s *Style) Tcell() tcell.Style {
	st := tcell.StyleDefault
	bg := s.GetBackground()
	fg := s.GetForeground()
	st = st.Background(tcell.NewRGBColor(bg.Red(), bg.Green(), bg.Blue()))
	st = st.Foreground(tcell.NewRGBColor(fg.Red(), fg.Green(), fg.Blue()))
	return st
}
