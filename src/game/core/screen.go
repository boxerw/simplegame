package core

import (
	"fmt"
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
}

func (screen *Screen) init() {
	if err := termbox.Init(); err != nil {
		panic(fmt.Sprintf("termbox init failed, %v", err))
	}

	termbox.Clear(termbox.RGBToAttribute(154, 22, 30), termbox.RGBToAttribute(0, 0, 0))
}

func (screen *Screen) shut() {
	termbox.Close()
}
