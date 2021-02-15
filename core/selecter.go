package core

type Selector struct {
    selectedComponent Selectable
}

func NewSelector() *Selector {
    return &Selector{}
}

func(s *Selector) GetSelected() Selectable {
    return s.selectedComponent
}

func(s *Selector) ToggleSelect(components []Component) bool {


    if s.selectedComponent != nil {
        if s.selectedComponent.ToggleSelect() {
            // the component handled the tab itself, and selected a child component
            return true
        }
    }

    var currentFound bool

    // we need to select the next top-level component
    for _, comp := range components {

        sel, ok := comp.(Selectable)
        if !ok {
            continue
        }

        // if nothing is selected, take the first component we see
        if s.selectedComponent == nil && sel.ToggleSelect() {
            s.selectedComponent = sel
            return true
        }

        // see if the selected component can itself select a new component
        if sel == s.selectedComponent {
            if sel.ToggleSelect() {
                return true
            }
            currentFound = true
        }

        if currentFound && sel.ToggleSelect() {
            s.selectedComponent = sel
            return true
        }
    }


    // we found the current selection but there was nothing after it in the list to select
    // let's go around again and find the first thing available for selection
    if currentFound {
        for _, comp := range components {
            sel, ok := comp.(Selectable)
            if !ok {
                continue
            }
            if sel.ToggleSelect() {
                s.selectedComponent = sel
                return true
            }
        }
    }

    // give up, we need to find something higher in the hierarchy to select
    return false
}
