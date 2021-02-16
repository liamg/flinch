package core

type Alignment uint8

const (
	AlignLeft Alignment = iota
	AlignRight
	AlignCenter
	AlignTop    = AlignLeft
	AlignBottom = AlignRight
)
