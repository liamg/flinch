package core

type Style interface {
	GetForeground() Colour
	GetBackground() Colour
	SetForeground(Colour)
	SetBackground(Colour)
}
