package core

type Justification uint8

const (
	JustifyLeft Justification = iota
	JustifyRight
	JustifyCenter
	JustifyFill
	JustifyTop    = JustifyLeft
	JustifyBottom = JustifyRight
)
