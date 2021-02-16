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
	return s.Select(components, true)
}

func (s *Selector) Deselect() {
	if s.selectedComponent != nil {
		s.selectedComponent.Deselect()
		s.selectedComponent = nil
	}
}

func (s *Selector) Select(components []Component, loop bool) bool {
	return s.toggleSelect(components, 1, loop)
}

func (s *Selector) toggleSelect(components []Component, inc int, loop bool) bool {

	var start int
	end := len(components)

	if inc < 0 {
		start = end
		end = 0
	}

	if s.selectedComponent != nil {
		if s.selectedComponent.Select(false) {
			// the component handled the tab itself, and selected a child component
			return true
		}
	}

	if !loop {

		var currentFound bool

		// we need to select the next top-level component
		for i := start; i != end; i += inc {

			comp := components[i]

			sel, ok := comp.(Selectable)
			if !ok {
				continue
			}

			// if nothing is selected, take the first component we see
			if s.selectedComponent == nil && sel.Select(false) {
				s.selectedComponent = sel
				return true
			}

			// see if the selected component can itself select a new component
			if sel == s.selectedComponent {
				currentFound = true
				continue
			}

			if currentFound && sel.Select(false) {
				s.selectedComponent.Deselect()
				s.selectedComponent = sel
				return true
			}
		}

	} else {

		// we found the current selection but there was nothing after it in the list to select
		// let's go around again and find the first thing available for selection
		for i := start; i != end; i += inc {
			comp := components[i]
			sel, ok := comp.(Selectable)
			if !ok {
				continue
			}
			if sel.Select(false) {
				if s.selectedComponent != nil {
					s.selectedComponent.Deselect()
				}
				s.selectedComponent = sel
				return true
			}
		}

	}

	// give up, we need to find something higher in the hierarchy to select
	return false
}
