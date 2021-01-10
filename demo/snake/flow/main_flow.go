package flow

import (
	"github.com/nsf/termbox-go"
	"simplegame/core"
	"simplegame/shell"
)

type MainFlow struct {
	shell.Scene
	shell.ControllerEventHook
	controller shell.Controller
	offset     shell.Posi2D
}

func (mainFlow *MainFlow) Init(object core.Object, name string) {
	scene, ok := object.(shell.Scene)
	if !ok {
		panic("not scene")
	}

	mainFlow.Scene = scene
	mainFlow.controller = shell.NewController(scene.GetEnvironment())
	mainFlow.controller.AddHook(mainFlow)
}

func (mainFlow *MainFlow) Shut() {
	mainFlow.controller.Destroy()
}

func (mainFlow *MainFlow) Update(frameCtx core.FrameContext) {
	mainFlow.controller.Update(frameCtx)

	var bitmaps shell.BitMaps
	bitmaps.Width = 30
	bitmaps.Matrix = make([]shell.Pixel, 30*bitmaps.Width)
	bitmaps.Matrix[0] = shell.Pixel{
		Ch: '0',
	}
	bitmaps.Matrix[29] = shell.Pixel{
		Ch: '1',
	}
	bitmaps.Matrix[29*bitmaps.Width] = shell.Pixel{
		Ch: '2',
	}
	bitmaps.Matrix[29*bitmaps.Width+29] = shell.Pixel{
		Ch: '3',
	}

	bitmaps.Origin.SetX(0)
	bitmaps.Origin.SetY(0)

	mainFlow.GetEnvironment().GetValue("screen").(shell.Screen).DrawMaps(0, mainFlow.offset, &bitmaps)
}

func (mainFlow *MainFlow) OnControllerKeyPress(controller shell.Controller, key termbox.Key, ch rune) bool {
	switch key {
	case termbox.KeyArrowUp:
		mainFlow.offset.SetY(mainFlow.offset.GetY() - 1)
	case termbox.KeyArrowDown:
		mainFlow.offset.SetY(mainFlow.offset.GetY() + 1)
	case termbox.KeyArrowLeft:
		mainFlow.offset.SetX(mainFlow.offset.GetX() - 1)
	case termbox.KeyArrowRight:
		mainFlow.offset.SetX(mainFlow.offset.GetX() + 1)
	}

	return true
}

func (mainFlow *MainFlow) OnControllerMousePress(controller shell.Controller, key termbox.Key, posi shell.Posi2D) bool {
	mainFlow.offset = posi

	return true
}
