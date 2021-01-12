package logic

import (
	"github.com/nsf/termbox-go"
	"simplegame/client"
	"simplegame/core"
)

type CoverStage struct {
	client.Scene
	screen client.Screen
}

func (coverStage *CoverStage) Init(object core.Object, name string) {
	coverStage.Scene = object.(client.Scene)
	coverStage.screen = coverStage.GetEnvironment().GetValue("screen").(client.Screen)
	coverStage.screen.SetCanvasFGBG(termbox.ColorWhite, termbox.ColorBlack)
}

func (coverStage *CoverStage) Shut() {
}

func (coverStage *CoverStage) Update(frameCtx core.FrameContext) {
	tips := "按'q'键退出，按其他任意键进入游戏"

	size := coverStage.screen.GetCanvasSize()
	pos := client.Posi2D{size.GetX()/2 - client.StringWidth(tips)/2, int(float32(size.GetY()) * 0.8)}

	coverStage.screen.DrawText(50, pos, tips, termbox.AttrBlink|termbox.ColorLightGray, termbox.ColorBlue)
}
