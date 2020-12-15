package logic

import (
	"fmt"
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
	screen     *Screen
	caches     []cache
	frameCount int
	beginTime  time.Time
}

func (logic *MainScreen) Name() string {
	return "MainScreen"
}

func (logic *MainScreen) Init(object Object) {
	logic.screen = object.(*Screen)
	logic.screen.GetContext().SetValue("MainScreen", logic.screen)
	logic.screen.SetCanvasFGBG(termbox.ColorRed, termbox.ColorWhite)
	rand.Seed(time.Now().UnixNano())
	logic.beginTime = time.Now()
}

func (logic *MainScreen) Shut() {
}

func (logic *MainScreen) Update() {
	if len(logic.caches) < 1024 {
		if logic.frameCount%3 == 0 {
			if len(logic.caches) <= 0 {
				w, h := termbox.Size()
				c := cache{
					x:     rand.Intn(w) / 2 * 2,
					y:     rand.Intn(h),
					color: termbox.Attribute(rand.Int()%256) + 1,
				}
				logic.caches = append(logic.caches, c)
			} else {
				c := cache{
					x:     logic.caches[len(logic.caches)-1].x + (rand.Intn(3)-1)*2,
					y:     logic.caches[len(logic.caches)-1].y + (rand.Intn(3) - 1),
					color: termbox.Attribute(rand.Int()%256) + 1,
				}
				w, h := termbox.Size()
				if c.x < 0 {
					c.x = w - c.x
				} else if c.x >= w {
					c.x = c.x - w
				}
				if c.y < 0 {
					c.y = h - 1
				} else if c.y >= h {
					c.y = 0
				}

				logic.caches = append(logic.caches, c)
			}
		}
	} else {
		logic.caches = nil
	}

	for _, c := range logic.caches {
		logic.screen.DrawText(1, c.x, c.y, termbox.ColorDefault, c.color, "  ")
	}

	logic.frameCount++

	dur := float64(time.Now().Sub(logic.beginTime) / time.Second)
	if dur <= 0 {
		dur = 1
	}
	frames := float64(logic.frameCount) / dur

	logic.screen.DrawText(100, 0, 0, termbox.ColorWhite, termbox.ColorRed, fmt.Sprintf("帧数：%.2f 长度：%d", frames, len(logic.caches)))
}
