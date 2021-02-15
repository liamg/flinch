package core

import "github.com/gdamore/tcell/v2"

type Selectable interface {
	Component
	ToggleSelect(loop bool) bool
	HandleKeypress(key *tcell.EventKey)
}
