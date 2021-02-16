package core

import "github.com/gdamore/tcell/v2"

type Selectable interface {
	Component
	Select() bool
	Deselect()
	HandleKeypress(key *tcell.EventKey)
}
