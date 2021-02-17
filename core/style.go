package core

import (
	"github.com/gdamore/tcell/v2"
)

type Style struct {
	inheritFg bool
	inheritBg bool
	invert    bool
	underline bool
	bold      bool
	bg        Colour
	fg        Colour
}

var StyleDefault = Style{
	bg: Colour([3]int32{0x0d, 0x19, 0x25}),
	fg: Colour([3]int32{0xd9, 0xe5, 0xf1}),
}

var StyleInherit = Style{
	inheritBg: true,
	inheritFg: true,
}

var StyleSelected = StyleDefault.Invert()

var StyleFaint = StyleDefault.SetForeground([3]int32{0x80, 0x80, 0x80})

var StyleButton = StyleFaint
var StyleButtonSelected = StyleDefault.Invert().Bold(true)

var StyleInput = StyleDefault

func (s Style) Invert() Style {
	s.invert = !s.invert
	return s
}

func (s Style) Underline(on bool) Style {
	s.underline = on
	return s
}

func (s Style) Bold(on bool) Style {
	s.bold = on
	return s
}

func (s Style) RemoveBackground() Style {
	s.inheritBg = true
	return s
}

func (s Style) GetForeground() Colour {
	return s.fg
}

func (s Style) GetBackground() Colour {
	return s.bg
}

func (s Style) SetForeground(colour Colour) Style {
	s.fg = colour
	return s
}

func (s Style) SetBackground(colour Colour) Style {
	s.bg = colour
	return s
}

func (s Style) Tcell() tcell.Style {
	st := tcell.StyleDefault
	if !s.inheritBg {
		bg := s.GetBackground()
		st = st.Background(tcell.NewRGBColor(bg.Red(), bg.Green(), bg.Blue()))
	}
	if !s.inheritFg {
		fg := s.GetForeground()
		st = st.Foreground(tcell.NewRGBColor(fg.Red(), fg.Green(), fg.Blue()))
	}
	return st.Reverse(s.invert).Underline(s.underline).Bold(s.bold)
}
