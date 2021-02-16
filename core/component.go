package core

type Component interface {
	// Render draws the component to the provided canvas
	Render(canvas Canvas)
	// MinimumSize returns the minimum size required by the component
	MinimumSize() Size

}
