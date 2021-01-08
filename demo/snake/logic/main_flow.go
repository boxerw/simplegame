package logic

import (
	"github.com/nsf/termbox-go"
	"simplegame/core"
	"simplegame/shell"
)

type MainFlow struct {
	shell.Scene
	shell.ControllerEventHook
	controller shell.Controller
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
}

func (mainFlow *MainFlow) OnControllerKeyPress(controller shell.Controller, key termbox.Key, ch rune) bool {
	if ch == 'q' {
		mainExecute := controller.GetEnvironment().GetValue("execute").(core.Execute)
		mainExecute.Shut()
	}

	return true
}
