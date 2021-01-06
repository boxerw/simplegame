package logic

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"simplegame/core"
	"simplegame/model"
	"time"
)

type ShowInfo struct {
	model.ControllerEventHook
	screen                       model.Screen
	controller                   model.Controller
	beginTime                    time.Time
	frameCount                   int
	keyPressInfo, mousePressInfo string
	hide                         bool
}

func (showInfo *ShowInfo) Init(object core.Object, name string) {
	screen, ok := object.(model.Screen)
	if !ok {
		panic("not screen")
	}

	showInfo.screen = screen
	showInfo.beginTime = time.Now()
	showInfo.hide = false
	showInfo.keyPressInfo = "键盘：[]"
	showInfo.mousePressInfo = "鼠标：[]"
	showInfo.controller = model.NewController(screen.GetEnvironment())
	showInfo.controller.AddHook(showInfo)
}

func (showInfo *ShowInfo) Shut() {
	showInfo.controller.Destroy()
}

func (showInfo *ShowInfo) Update() {
	showInfo.frameCount++

	dur := float64(time.Now().Sub(showInfo.beginTime) / time.Second)
	if dur <= 0 {
		dur = 1
	}
	frames := float64(showInfo.frameCount) / dur

	if !showInfo.hide {
		canvasSize := showInfo.screen.GetCanvasSize()
		showInfo.screen.DrawText(100, model.Posi2D{0, 0},
			fmt.Sprintf("屏幕：[W:%d H:%d]，帧：[%.2f/s]", canvasSize.GetX(), canvasSize.GetY(), frames),
			termbox.ColorWhite, termbox.ColorRed)
	}

	showInfo.controller.Update()

	if !showInfo.hide {
		showInfo.screen.DrawText(100, model.Posi2D{0, 1},
			showInfo.keyPressInfo,
			termbox.ColorWhite, termbox.ColorRed)

		showInfo.screen.DrawText(100, model.Posi2D{0, 2},
			showInfo.mousePressInfo,
			termbox.ColorWhite, termbox.ColorRed)
	}
}

func (showInfo *ShowInfo) OnControllerKeyPress(controller model.Controller, key termbox.Key, ch rune) bool {
	if key == termbox.KeyF1 {
		showInfo.hide = !showInfo.hide
	}

	showInfo.keyPressInfo = fmt.Sprintf("键盘：[KEY:%v CH:%v]", key, ch)

	return true
}

func (showInfo *ShowInfo) OnControllerMousePress(controller model.Controller, key termbox.Key, posi model.Posi2D) bool {
	showInfo.mousePressInfo = fmt.Sprintf("鼠标：[KEY:%v X:%d Y:%d]", key, posi.GetX(), posi.GetY())

	return true
}
