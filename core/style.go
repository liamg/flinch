package core

import (
	"github.com/gdamore/tcell/v2"
)

type Style struct {
	InheritFg   bool             `mapstructure:"inherit_fg"`
	InheritBg   bool             `mapstructure:"inherit_bg"`
	Invert      bool             `mapstructure:"invert"`
	Underline   bool             `mapstructure:"underline"`
	Bold        bool             `mapstructure:"bold"`
	Bg          Colour           `mapstructure:"bg"`
	Fg          Colour           `mapstructure:"fg"`
	ButtonEdges ButtonEdgesStyle `mapstructure:"button_edges"`
}

type ButtonEdgesStyle string

const (
	ButtonEdgesStyleAngled  ButtonEdgesStyle = "angled"
	ButtonEdgesStyleFlat    ButtonEdgesStyle = "flat"
	ButtonEdgesStyleRound   ButtonEdgesStyle = "round"
	ButtonEdgesStyleSlanted ButtonEdgesStyle = "slanted"
	ButtonEdgesStyleNone    ButtonEdgesStyle = "none"
	ButtonEdgesStyleDefault ButtonEdgesStyle = ButtonEdgesStyleAngled
)

var (
	StyleDefault        Style
	StyleSelected       Style
	StyleFaint          Style
	StyleButton         Style
	StyleButtonSelected Style
	StyleInput          Style
)

var (
	buttonEdges = map[ButtonEdgesStyle][2]rune{
		ButtonEdgesStyleAngled:  {'', ''},
		ButtonEdgesStyleFlat:    {'█', '█'},
		ButtonEdgesStyleRound:   {'', ''},
		ButtonEdgesStyleSlanted: {'', ''},
		ButtonEdgesStyleNone:    {' ', ' '},
	}
)

var StyleInherit = Style{
	InheritBg:   true,
	InheritFg:   true,
	ButtonEdges: ButtonEdgesStyleDefault,
}

func SetStyle(style Style) {
	StyleDefault = style
	StyleInput = StyleDefault
	StyleSelected = StyleDefault.ToggleInvert()
	StyleFaint = StyleDefault.
		SetInheritForeground(false).
		SetForeground(FaintColour(StyleDefault.Fg))
	StyleButton = StyleFaint
	StyleButtonSelected = StyleDefault.ToggleInvert().SetBold(true)
}

func (s Style) ToggleInvert() Style {
	s.Invert = !s.Invert
	return s
}

func (s Style) SetUnderline(on bool) Style {
	s.Underline = on
	return s
}

func (s Style) SetBold(on bool) Style {
	s.Bold = on
	return s
}

func (s Style) RemoveBackground() Style {
	s.InheritBg = true
	return s
}

func (s Style) GetForeground() Colour {
	return s.Fg
}

func (s Style) GetBackground() Colour {
	return s.Bg
}

func (s Style) SetForeground(colour Colour) Style {
	s.Fg = colour
	return s
}

func (s Style) SetBackground(colour Colour) Style {
	s.Bg = colour
	return s
}

func (s Style) SetInheritForeground(inheritFg bool) Style {
	s.InheritFg = inheritFg
	return s
}

func (s Style) SetInheritBackground(inheritBg bool) Style {
	s.InheritBg = inheritBg
	return s
}

func (s Style) SetButtonEdges(b ButtonEdgesStyle) Style {
	s.ButtonEdges = b
	return s
}

func (s Style) GetButtonEdges() [2]rune {
	edges, ok := buttonEdges[s.ButtonEdges]
	if !ok {
		edges = buttonEdges[ButtonEdgesStyleDefault]
	}

	return edges
}

func (s Style) Tcell() tcell.Style {
	st := tcell.StyleDefault
	if !s.InheritBg {
		bg := s.GetBackground()
		st = st.Background(tcell.NewRGBColor(bg.Red(), bg.Green(), bg.Blue()))
	}
	if !s.InheritFg {
		fg := s.GetForeground()
		st = st.Foreground(tcell.NewRGBColor(fg.Red(), fg.Green(), fg.Blue()))
	}
	return st.Reverse(s.Invert).Underline(s.Underline).Bold(s.Bold)
}
