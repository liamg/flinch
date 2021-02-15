package core

type Component interface {
	// Render draws the component to the provided canvas
	Render(canvas Canvas)
	// Size returns the minimum size required by the component
	Size(parent Canvas) (int, int)
}
