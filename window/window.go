package window

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
	WithJustification(j core.Justification) Window
	Show() error
	Close()
}

type window struct {
	container        core.Container
	mu               sync.Mutex
	screen           tcell.Screen
	focusedComponent core.Component
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

func (w *window) WithJustification(j core.Justification) Window {
	w.container.WithJustification(j)
	return w
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
	return nil
}

func (w *window) Show() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	for {

		w.screen.Clear()
		canvas := NewBaseCanvas(w.screen)

		w.container.Render(canvas)
		//w.screen.SetCell(5, 5, tcell.StyleDefault, '!')
		w.screen.Show()

		switch ev := w.screen.PollEvent().(type) {
		case *tcell.EventResize:
			w.screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape { // TODO remove this - just here to exit during testing
				w.Close()
				return nil
			}
		}

	}
}

func (w *window) Close() {
	w.screen.Fini()
}
