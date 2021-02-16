package core

import "github.com/gdamore/tcell/v2"

type Selectable interface {
	Component
	Select(loop bool) bool
	Deselect()
	HandleKeypress(key *tcell.EventKey)
}
