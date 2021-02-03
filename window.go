package flinch

import (
	"sync"

	"github.com/liamg/flinch/components"

	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type Window interface {
	// Size returns the minimum size required by the component
	Size() (int, int)
	Add(c core.Component)
	Remove(c core.Component)
	SetJustification(j core.Justification)
	Show() error
	Close()
}

type window struct {
	container        core.Container
	mu               sync.Mutex
	screen           tcell.Screen
	focusedComponent core.Component
	canvas           core.Canvas
}

func New() (Window, error) {
	w := &window{
		container: components.NewColumnLayout(),
	}
	if err := w.init(); err != nil {
		return nil, err
	}
	return w, nil
}

func (w *window) SetContainer(c core.Container) {
	w.container = c
}

func (w *window) SetJustification(j core.Justification) {
	w.container.SetJustification(j)
}

func (w *window) Size() (int, int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.screen == nil {
		return 0, 0
	}
	return w.screen.Size()
}

func (w *window) Add(component core.Component) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.focusedComponent == nil {
		w.focusedComponent = component
	}
	w.container.Add(component)
}

func (w *window) Remove(component core.Component) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.container.Remove(component)
}

func (w *window) Focus(c core.Component) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.focusedComponent = c
}

func (w *window) init() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	if err := screen.Init(); err != nil {
		return err
	}
	w.screen = screen
	w.canvas = NewBaseCanvas(w.screen)
	w.screen.Clear()
	return nil
}

func (w *window) Show() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.container.Render(w.canvas)
	//w.screen.SetCell(5, 5, tcell.StyleDefault, '!')
	w.screen.Show()

	// TODO add for loop here until exit
	<-(make(chan bool))

	return nil
}

func (w *window) Close() {
	w.screen.Fini()
}
