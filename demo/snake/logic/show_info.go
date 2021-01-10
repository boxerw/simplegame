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
	drawBeginTime                time.Time
	drawConsumeTime              time.Duration
	statDrawConsumeBeginTime     time.Time
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
	showInfo.statDrawConsumeBeginTime = time.Now()
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

func (showInfo *ShowInfo) OnBeginDrawing(screen shell.Screen, drawCache shell.DrawCache) bool {
	if showInfo.frameCtx == nil {
		return true
	}

	if !showInfo.hide {
		canvasSize := showInfo.GetCanvasSize()
		showInfo.DrawText(100, shell.Posi2D{0, 0},
			fmt.Sprintf("屏幕：[W:%d H:%d]，提交：[%d]，耗时：[%dms]，FPS：[%.2f/s]",
				canvasSize.GetX(),
				canvasSize.GetY(),
				drawCache.Size(),
				showInfo.drawConsumeTime.Milliseconds(),
				showInfo.frameCtx.GetFPS()),
			termbox.ColorWhite, termbox.ColorRed)

		showInfo.DrawText(100, shell.Posi2D{0, 1},
			showInfo.keyPressInfo,
			termbox.ColorWhite, termbox.ColorRed)

		showInfo.DrawText(100, shell.Posi2D{0, 2},
			showInfo.mousePressInfo,
			termbox.ColorWhite, termbox.ColorRed)
	}

	showInfo.drawBeginTime = time.Now()

	return true
}

func (showInfo *ShowInfo) OnEndDrawing(screen shell.Screen) bool {
	now := time.Now()

	delta := now.Sub(showInfo.drawBeginTime)
	if delta > showInfo.drawConsumeTime {
		showInfo.drawConsumeTime = delta
	}

	if now.Sub(showInfo.statDrawConsumeBeginTime).Seconds() > 3 {
		showInfo.statDrawConsumeBeginTime = now
		showInfo.drawConsumeTime = 0
	}

	return true
}
