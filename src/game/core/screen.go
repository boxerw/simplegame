package core

import (
	"fmt"
	"github.com/mattn/go-runewidth"
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
	drawCount int
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

	if screen.drawCount > 0 {
		termbox.Flush()
		screen.drawCount = 0
	}
}

func (screen *Screen) init() {
	if err := termbox.Init(); err != nil {
		panic(fmt.Sprintf("termbox Init failed, %v", err))
	}

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.SetOutputMode(termbox.Output256)
}

func (screen *Screen) shut() {
	termbox.Close()
}

func (screen *Screen) Draw(x, y int, fg, bg termbox.Attribute, text string) {
	if len(text) <= 0 {
		return
	}

	for _, c := range text {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}

	screen.drawCount++
}

func (screen *Screen) Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}
