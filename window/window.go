package window

import (
	"sync"
	"time"

	"github.com/liamg/flinch/components"

	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type Window interface {
	// Size returns the minimum size required by the component
	Size() (int, int)
	Add(c core.Component)
	Remove(c core.Component)
	SetSimulation(s bool)
	SetAlignment(j core.Alignment) Window
	Show() error
	OnKeypress(handler func(key *tcell.EventKey) bool)
	Close()
}

type window struct {
	container    core.Container
	mu           sync.Mutex
	screen       tcell.Screen
	keyHandlers  []func(key *tcell.EventKey) bool
	shouldClose  bool
	isSimulation bool
	visible      bool
}

type WindowOption func(w Window) error

func WindowOptionSimulation() WindowOption {
	return func(w Window) error {
		w.SetSimulation(true)
		return nil
	}
}

func New(opts ...WindowOption) (Window, error) {
	w := &window{
		container: components.NewColumnLayout(),
	}

	for _, opt := range opts {
		if err := opt(w); err != nil {
			return nil, err
		}
	}

	if err := w.init(); err != nil {
		return nil, err
	}
	return w, nil
}

func (w *window) SetSimulation(s bool) {
	w.isSimulation = s
}

func (w *window) SetContainer(c core.Container) {
	w.container = c
}

func (w *window) SetAlignment(j core.Alignment) Window {
	w.container.SetAlignment(j)
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
	var screen tcell.Screen
	var err error
	if w.isSimulation {
		screen = tcell.NewSimulationScreen("")
	} else {
		screen, err = tcell.NewScreen()
		if err != nil {
			return err
		}
	}

	if err = screen.Init(); err != nil {
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
		if !sel.Select() {
			sel.Deselect()
			sel.Select()
		}
	}
}

func (w *window) render() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.screen.Clear()
	w.screen.SetStyle(core.StyleDefault.Tcell())
	canvas := NewBaseCanvas(w.screen)
	canvas.Fill(' ', core.StyleDefault)
	w.container.Render(canvas)
	w.screen.Show()
}

func (w *window) sync() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.screen.Sync()
}

func (w *window) Show() error {

	w.mu.Lock()
	w.visible = true
	w.mu.Unlock()

	if sel, ok := w.container.(core.Selectable); ok {
		sel.Select()
	}

	go func() {
		// avoid stdout race from caller
		time.Sleep(time.Millisecond * 250)
		w.sync()
	}()

	for {

		w.render()

		switch ev := w.screen.PollEvent().(type) {
		case *tcell.EventResize:
			continue
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyTab:
				w.selectNext()
			case tcell.KeyEscape:
				w.Close()
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

		if w.shouldClose {
			w.mu.Lock()
			if w.visible {
				w.screen.Fini()
				w.visible = false
			}
			w.mu.Unlock()
			break
		}

	}

	return nil
}

func (w *window) Close() {
	w.mu.Lock()
	if !w.visible {
		w.screen.Fini()
	}
	w.mu.Unlock()
	w.shouldClose = true
}
