package logic

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	. "simple/game/core"
	"time"
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
	rand.Seed(time.Now().UnixNano())
}

func (logic *MainScreen) Shut() {
}

func (logic *MainScreen) Update() {
	if len(logic.caches) < 1024 {
		if len(logic.caches) <= 0 {
			w, h := termbox.Size()
			c := cache{
				x:     rand.Intn(w) / 2 * 2,
				y:     rand.Intn(h),
				color: termbox.Attribute(rand.Int()%256) + 1,
			}
			logic.caches = append(logic.caches, c)
		} else {
			randFun := func() int {
				v := rand.Intn(2)
				if v <= 0 {
					return -1
				}
				return v
			}

			c := cache{
				x:     logic.caches[len(logic.caches)-1].x + randFun(),
				y:     logic.caches[len(logic.caches)-1].y + randFun(),
				color: termbox.Attribute(rand.Int()%256) + 1,
			}
			w, h := termbox.Size()
			if c.x < 0 {
				c.x = w - 1
			} else if c.x >= w {
				c.x = 0
			}
			if c.y < 0 {
				c.y = h - 1
			} else if c.y >= h {
				c.y = 0
			}

			c.x = c.x / 2 * 2

			logic.caches = append(logic.caches, c)
		}
	} else {
		logic.caches = nil
	}

	for _, c := range logic.caches {
		logic.screen.DrawText(0, c.x, c.y, termbox.ColorDefault, c.color, "  ")
	}
}
