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
	container         core.Container
	mu                sync.Mutex
	screen            tcell.Screen
	keyHandlers       []func(key *tcell.EventKey) bool
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
	w.container.Add(component)
}

func (w *window) Remove(component core.Component) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.container.Remove(component)
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

func (w *window) OnKeypress(handler func(key *tcell.EventKey) bool) {
	w.keyHandlers = append(w.keyHandlers, handler)
}

func (w *window) selectNext() {
	if sel, ok := w.container.(core.Selectable); ok {
		sel.ToggleSelect()
	}
}

func (w *window) Show() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if sel, ok := w.container.(core.Selectable); ok {
		sel.ToggleSelect()
	}

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
			switch ev.Key() {
			case tcell.KeyTab:
				w.selectNext()
			case tcell.KeyEscape:
				w.Close()
				return nil
			default:
				var handled bool
				for _, handler := range w.keyHandlers {
					if handler(ev) {
						handled = true
						break
					}
				}
				if !handled {
					if sel, ok := w.container.(core.Selectable); ok {
						sel.HandleKeypress(ev)
					}
				}
			}
		}

	}
}

func (w *window) Close() {
	w.screen.Fini()
}
