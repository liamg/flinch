package core

import "github.com/gdamore/tcell/v2"

type Selectable interface {
    Component
    ToggleSelect() bool
    HandleKeypress(key *tcell.EventKey)
}
