package logic

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	. "simple/game/core"
)

type MainScreen struct {
	screen *Screen
}

func (logic *MainScreen) Name() string {
	return "MainScreen"
}

func (logic *MainScreen) Init(object Object) {
	logic.screen = object.(*Screen)
	logic.screen.GetContext().SetValue("MainScreen", logic.screen)
}

func (logic *MainScreen) Shut() {
}

func (logic *MainScreen) Update() {
	w, h := termbox.Size()
	x := rand.Intn(w) / 2 * 2
	y := rand.Intn(h)

	logic.screen.Draw(x, y, termbox.ColorDefault, termbox.Attribute(rand.Int()%256)+1, "  ")
}
