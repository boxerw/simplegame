package shell

import (
	"github.com/nsf/termbox-go"
	"simplegame/core"
	"unsafe"
)

type Screen interface {
	core.Object
	core.Event
	SetCanvasFGBG(fg, bg termbox.Attribute)
	GetCanvasFGBG() (fg, bg termbox.Attribute)
	GetCanvasSize() Posi2D
	DrawMaps(layer int, posi Posi2D, maps Maps)
	DrawText(layer int, posi Posi2D, text string, fg, bg termbox.Attribute)
	Flush()
}

func NewScreen(env core.Environment, components ...core.ComponentBundle) Screen {
	screen := &_Screen{}
	screen.Init(env, screen, func() {
		if err := termboxEx.Init(); err != nil {
			panic(err)
		}

		termbox.SetOutputMode(termbox.Output256)
		termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

		w, h := termbox.Size()
		screen.canvasSize = Posi2D{w, h}

		screen.inputChan = make(chan termbox.Event, 100)
		termboxEx.AddInputEventHook(unsafe.Pointer(screen), screen.inputChan)

	}, components...)
	return screen
}

type _Screen struct {
	core.ObjectModule
	core.EventModule
	drawCache          DrawCache
	canvasFG, canvasBG termbox.Attribute
	canvasSize         Posi2D
	inputChan          chan termbox.Event
}

func (screen *_Screen) Destroy() {
	screen.Shut(func() {
		termboxEx.RemoveInputEventHook(unsafe.Pointer(screen))
		close(screen.inputChan)

		termboxEx.Shut()
	})
}

func (screen *_Screen) Update(frameCtx core.FrameContext) {
	if screen.GetDisabled() {
		return
	}

	screen.ObjectModule.Update(frameCtx)

	func() {
		for {
			select {
			case ev := <-screen.inputChan:
				switch ev.Type {
				case termbox.EventResize:
					oldSize := screen.canvasSize
					screen.canvasSize = Posi2D{ev.Width, ev.Height}

					screen.RangeHooks(func(hook core.Hook) bool {
						return screen.ExecFunc(func() bool {
							if hook, ok := hook.(ScreenEvent); ok {
								return hook.OnScreenCanvasSizeChange(screen, oldSize, screen.canvasSize)
							}
							return true
						})
					})
				}
			default:
				return
			}
		}
	}()

	screen.RangeHooks(func(hook core.Hook) bool {
		return screen.ExecFunc(func() bool {
			if hook, ok := hook.(ScreenEvent); ok {
				return hook.OnBeginDrawing(screen, screen.drawCache)
			}
			return true
		})
	})

	screen.drawCache.Drawing(screen.canvasFG, screen.canvasBG)

	screen.RangeHooks(func(hook core.Hook) bool {
		return screen.ExecFunc(func() bool {
			if hook, ok := hook.(ScreenEvent); ok {
				return hook.OnEndDrawing(screen)
			}
			return true
		})
	})
}

func (screen *_Screen) SetCanvasFGBG(fg, bg termbox.Attribute) {
	if screen.GetDestroyed() {
		return
	}

	screen.canvasFG, screen.canvasBG = fg, bg
}

func (screen *_Screen) GetCanvasFGBG() (fg, bg termbox.Attribute) {
	return screen.canvasFG, screen.canvasBG
}

func (screen *_Screen) GetCanvasSize() Posi2D {
	return screen.canvasSize
}

func (screen *_Screen) DrawMaps(layer int, posi Posi2D, maps Maps) {
	if screen.GetDestroyed() || maps == nil {
		return
	}

	screen.drawCache.AddItem(&DrawItem{
		Layer: layer,
		Posi:  posi,
		Maps:  maps,
	})
}

func (screen *_Screen) DrawText(layer int, posi Posi2D, text string, fg, bg termbox.Attribute) {
	if screen.GetDestroyed() || len(text) <= 0 {
		return
	}

	screen.drawCache.AddItem(&DrawItem{
		Layer: layer,
		Posi:  posi,
		Text: &Text{
			Content: text,
			Fg:      fg,
			Bg:      bg,
		},
	})
}

func (screen *_Screen) Flush() {
	if screen.GetDestroyed() {
		return
	}

	screen.RangeHooks(func(hook core.Hook) bool {
		return screen.ExecFunc(func() bool {
			if hook, ok := hook.(ScreenEvent); ok {
				return hook.OnBeginDrawing(screen, screen.drawCache)
			}
			return true
		})
	})

	screen.drawCache.Drawing(screen.canvasFG, screen.canvasBG)

	screen.RangeHooks(func(hook core.Hook) bool {
		return screen.ExecFunc(func() bool {
			if hook, ok := hook.(ScreenEvent); ok {
				return hook.OnEndDrawing(screen)
			}
			return true
		})
	})
}
