package logic

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"simplegame/core"
	"simplegame/shell"
)

type ShowInfo struct {
	shell.Screen
	shell.ControllerEventHook
	controller                   shell.Controller
	keyPressInfo, mousePressInfo string
	hide                         bool
}

func (showInfo *ShowInfo) Init(object core.Object, name string) {
	screen, ok := object.(shell.Screen)
	if !ok {
		panic("not screen")
	}

	showInfo.Screen = screen
	showInfo.hide = false
	showInfo.keyPressInfo = "键盘：[]"
	showInfo.mousePressInfo = "鼠标：[]"
	showInfo.controller = shell.NewController(screen.GetEnvironment())
	showInfo.controller.AddHook(showInfo)
}

func (showInfo *ShowInfo) Shut() {
	showInfo.controller.Destroy()
}

func (showInfo *ShowInfo) Update(frameCtx core.FrameContext) {
	if !showInfo.hide {
		canvasSize := showInfo.GetCanvasSize()
		showInfo.DrawText(100, shell.Posi2D{0, 0},
			fmt.Sprintf("屏幕：[W:%d H:%d]，FPS：[%.2f/s]", canvasSize.GetX(), canvasSize.GetY(), frameCtx.GetFPS()),
			termbox.ColorWhite, termbox.ColorRed)
	}

	showInfo.controller.Update(frameCtx)

	if !showInfo.hide {
		showInfo.DrawText(100, shell.Posi2D{0, 1},
			showInfo.keyPressInfo,
			termbox.ColorWhite, termbox.ColorRed)

		showInfo.DrawText(100, shell.Posi2D{0, 2},
			showInfo.mousePressInfo,
			termbox.ColorWhite, termbox.ColorRed)
	}
}

func (showInfo *ShowInfo) OnControllerKeyPress(controller shell.Controller, key termbox.Key, ch rune) bool {
	if key == termbox.KeyF1 {
		showInfo.hide = !showInfo.hide
	}

	showInfo.keyPressInfo = fmt.Sprintf("键盘：[KEY:%v CH:%v]", key, ch)

	return true
}

func (showInfo *ShowInfo) OnControllerMousePress(controller shell.Controller, key termbox.Key, posi shell.Posi2D) bool {
	showInfo.mousePressInfo = fmt.Sprintf("鼠标：[KEY:%v X:%d Y:%d]", key, posi.GetX(), posi.GetY())

	return true
}
