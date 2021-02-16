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
	return s.toggleSelect(components, -1)
}

func (s *Selector) Next(components []Component) bool {
	return s.Select(components)
}

func (s *Selector) Deselect() {
	if s.selectedComponent != nil {
		s.selectedComponent.Deselect()
		s.selectedComponent = nil
	}
}

func (s *Selector) Select(components []Component) bool {
	return s.toggleSelect(components, 1)
}

func (s *Selector) toggleSelect(components []Component, inc int) bool {

	var start int
	end := len(components)

	if inc < 0 {
		start = end
		end = 0
	}

	if s.selectedComponent != nil {
		if s.selectedComponent.Select() {
			// the component handled the tab itself, and selected a child component
			return true
		}
	}

	var currentFound bool

	// we need to select the next top-level component
	for i := start; i != end; i += inc {

		comp := components[i]

		sel, ok := comp.(Selectable)
		if !ok {
			continue
		}

		// if nothing is selected, take the first component we see
		if s.selectedComponent == nil && sel.Select() {
			s.selectedComponent = sel
			return true
		}

		// see if the selected component can itself select a new component
		if sel == s.selectedComponent {
			currentFound = true
			continue
		}

		if currentFound && sel.Select() {
			s.selectedComponent.Deselect()
			s.selectedComponent = sel
			return true
		}
	}

	// give up, we need to find something higher in the hierarchy to select
	return false
}
