package logic

import (
	"github.com/nsf/termbox-go"
	"simplegame/core"
	"simplegame/shell"
)

type CoverStage struct {
	shell.Scene
	screen shell.Screen
}

func (coverStage *CoverStage) Init(object core.Object, name string) {
	coverStage.Scene = object.(shell.Scene)
	coverStage.screen = coverStage.GetEnvironment().GetValue("screen").(shell.Screen)
}

func (coverStage *CoverStage) Shut() {
}

func (coverStage *CoverStage) Update(frameCtx core.FrameContext) {
	tips := "按'q'键退出，按其他任意键进入游戏"

	size := coverStage.screen.GetCanvasSize()
	pos := shell.Posi2D{size.GetX()/2 - shell.StringWidth(tips)/2, int(float32(size.GetY()) * 0.8)}

	coverStage.screen.DrawText(50, pos, tips, termbox.AttrBlink|termbox.ColorLightGray, termbox.ColorBlue)
}
