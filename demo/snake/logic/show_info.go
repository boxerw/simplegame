package logic

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"simplegame/core"
	"simplegame/shell"
	"time"
)

type ShowInfo struct {
	shell.Screen
	shell.ControllerEventHook
	shell.ScreenEventHook
	core.HookModule
	controller                   shell.Controller
	keyPressInfo, mousePressInfo string
	hide                         bool
	frameCtx                     core.FrameContext
	statBeginTime                time.Time
	statDrawBeginTime            time.Time
	statDrawConsumeTime          time.Duration
	statDrawCommit               int
}

func (showInfo *ShowInfo) Init(object core.Object, name string) {
	screen, ok := object.(shell.Screen)
	if !ok {
		panic("not screen")
	}

	showInfo.Screen = screen
	showInfo.Screen.AddHook(showInfo)
	showInfo.hide = false
	showInfo.keyPressInfo = "键盘：[]"
	showInfo.mousePressInfo = "鼠标：[]"
	showInfo.controller = shell.NewController(screen.GetEnvironment())
	showInfo.controller.AddHook(showInfo)
	showInfo.statBeginTime = time.Now()
	showInfo.statDrawCommit = 0
}

func (showInfo *ShowInfo) Shut() {
	showInfo.controller.Destroy()
}

func (showInfo *ShowInfo) Update(frameCtx core.FrameContext) {
	showInfo.controller.Update(frameCtx)
	showInfo.frameCtx = frameCtx
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

func (showInfo *ShowInfo) OnBeginDrawing(screen shell.Screen) bool {
	if showInfo.frameCtx == nil {
		return true
	}

	if showInfo.statDrawCommit < screen.GetDrawCache().Size() {
		showInfo.statDrawCommit = screen.GetDrawCache().Size()
	}

	if !showInfo.hide {
		canvasSize := showInfo.GetCanvasSize()
		showInfo.DrawText(100, shell.Posi2D{0, 0},
			fmt.Sprintf("屏幕：[W:%d H:%d FPS：%.2f/s]",
				canvasSize.GetX(),
				canvasSize.GetY(),
				showInfo.frameCtx.GetFPS()),
			termbox.ColorWhite, termbox.ColorRed)

		showInfo.DrawText(100, shell.Posi2D{0, 1},
			fmt.Sprintf("渲染：[提交：%d 消耗：%dms]",
				showInfo.statDrawCommit,
				showInfo.statDrawConsumeTime.Milliseconds()),
			termbox.ColorWhite, termbox.ColorRed)

		showInfo.DrawText(100, shell.Posi2D{0, 2},
			showInfo.keyPressInfo,
			termbox.ColorWhite, termbox.ColorRed)

		showInfo.DrawText(100, shell.Posi2D{0, 3},
			showInfo.mousePressInfo,
			termbox.ColorWhite, termbox.ColorRed)
	}

	showInfo.statDrawBeginTime = time.Now()

	return true
}

func (showInfo *ShowInfo) OnEndDrawing(screen shell.Screen) bool {
	now := time.Now()

	delta := now.Sub(showInfo.statDrawBeginTime)
	if delta > showInfo.statDrawConsumeTime {
		showInfo.statDrawConsumeTime = delta
	}

	if now.Sub(showInfo.statBeginTime).Seconds() > 1 {
		showInfo.statBeginTime = now
		showInfo.statDrawConsumeTime = 0
		showInfo.statDrawCommit = 0
	}

	return true
}
