package core

type Selector struct {
	selectedComponent Selectable
}

func NewSelector() *Selector {
	return &Selector{}
}

func (s *Selector) GetSelected() Selectable {
	return s.selectedComponent
}

func (s *Selector) Previous(components []Component) bool {
	return s.toggleSelect(components, -1, true)
}

func (s *Selector) Next(components []Component) bool {
	return s.ToggleSelect(components, true)
}

func (s *Selector) ToggleSelect(components []Component, loop bool) bool {
	return s.toggleSelect(components, 1, loop)
}

func (s *Selector) toggleSelect(components []Component, inc int, loop bool) bool {

	if s.selectedComponent != nil {
		if s.selectedComponent.ToggleSelect(false) {
			// the component handled the tab itself, and selected a child component
			if loop {
				panic(s.selectedComponent)
			}
			return true
		}
	}

	var currentFound bool

	var start int
	end := len(components)

	// we need to select the next top-level component
	for i := start; i != end; i += inc {

		comp := components[i]

		sel, ok := comp.(Selectable)
		if !ok {
			continue
		}

		// if nothing is selected, take the first component we see
		if s.selectedComponent == nil && sel.ToggleSelect(false) {
			s.selectedComponent = sel
			return true
		}

		// see if the selected component can itself select a new component
		if sel == s.selectedComponent {
			currentFound = true
			continue
		}

		if currentFound && sel.ToggleSelect(false) {
			s.selectedComponent = sel
			return true
		}
	}

	// we found the current selection but there was nothing after it in the list to select
	// let's go around again and find the first thing available for selection
	if loop && currentFound {
		for i := start; i != end; i += inc {
			comp := components[i]
			sel, ok := comp.(Selectable)
			if !ok {
				continue
			}
			if sel.ToggleSelect(false) {
				s.selectedComponent = sel
				return true
			}
		}
	}

	// give up, we need to find something higher in the hierarchy to select
	return false
}
