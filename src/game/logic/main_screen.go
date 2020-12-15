package logic

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	. "simple/game/core"
)

type cache struct {
	x, y  int
	color termbox.Attribute
}

type MainScreen struct {
	screen *Screen
	caches []cache
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
	if len(logic.caches) < 300 {
		w, h := termbox.Size()
		c := cache{
			x:     rand.Intn(w) / 2 * 2,
			y:     rand.Intn(h),
			color: termbox.Attribute(rand.Int()%256) + 1,
		}
		logic.caches = append(logic.caches, c)
	}

	for _, c := range logic.caches {
		logic.screen.DrawText(0, c.x, c.y, termbox.ColorDefault, c.color, "  ")
	}
}
