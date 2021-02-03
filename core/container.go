package core

type Container interface {
	Component
	Add(c Component)
	Remove(c Component)
	SetJustification(justification Justification)
}
