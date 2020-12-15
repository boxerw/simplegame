package core

import (
	"github.com/nsf/termbox-go"
)

func NewScreen(ctx *Context, logic ...Logic) *Screen {
	screen := &Screen{}
	screen.initObject(ctx, screen, logic...)
	screen.init()
	return screen
}

type Screen struct {
	_ObjectBase
	drawCache          _DrawCache
	canvasFG, canvasBG termbox.Attribute
}

func (screen *Screen) Destroy() {
	screen.shut()
	screen.shutObject()
}

func (screen *Screen) Update() {
	if screen.GetZombie() {
		return
	}

	screen.updateObject()

	screen.drawCache.Drawing(screen.canvasFG, screen.canvasBG)
}

func (screen *Screen) init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.SetOutputMode(termbox.Output256)
}

func (screen *Screen) shut() {
	termbox.Close()
}

func (screen *Screen) SetCanvasFGBG(fg, bg termbox.Attribute) {
	if screen.GetZombie() {
		return
	}

	screen.canvasFG, screen.canvasBG = fg, bg
}

func (screen *Screen) GetCanvasFGBG() (fg, bg termbox.Attribute) {
	return screen.canvasFG, screen.canvasBG
}

func (screen *Screen) DrawMaps(layer, x, y int, maps Maps) {
	if screen.GetZombie() || len(maps) <= 0 {
		return
	}

	screen.drawCache.AddItem(_DrawItem{
		Layer: layer,
		X:     x,
		Y:     y,
		Maps:  maps,
	})
}

func (screen *Screen) DrawShape(layer, x, y int, fg, bg termbox.Attribute, shape Shape) {
	if screen.GetZombie() || len(shape) <= 0 {
		return
	}

	maps := make(Maps, len(shape))
	for i := 0; i < len(shape); i++ {
		maps[i] = make([]Pixel, len(shape[i]))
		for j := 0; j < len(shape[i]); j++ {
			maps[i][j] = Pixel{
				FG:   fg,
				BG:   bg,
				Char: shape[i][j],
			}
		}
	}

	screen.DrawMaps(layer, x, y, maps)
}

func (screen *Screen) DrawText(layer, x, y int, fg, bg termbox.Attribute, text string) {
	if screen.GetZombie() || len(text) <= 0 {
		return
	}

	maps := make(Maps, 1)
	maps[0] = make([]Pixel, len(text))
	for j, c := range text {
		maps[0][j] = Pixel{
			FG:   fg,
			BG:   bg,
			Char: c,
		}
	}

	screen.DrawMaps(layer, x, y, maps)
}
