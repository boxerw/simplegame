package logic

import (
	"github.com/nsf/termbox-go"
	"simplegame/core"
	"simplegame/model"
)

type MainFlow struct {
	model.ControllerEventHook
	scene      model.Scene
	controller model.Controller
}

func (mainFlow *MainFlow) Init(object core.Object, name string) {
	scene, ok := object.(model.Scene)
	if !ok {
		panic("not scene")
	}

	mainFlow.scene = scene
	mainFlow.controller = model.NewController(scene.GetEnvironment())
	mainFlow.controller.AddHook(mainFlow)
}

func (mainFlow *MainFlow) Shut() {
	mainFlow.controller.Destroy()
}

func (mainFlow *MainFlow) Update() {
	mainFlow.controller.Update()
}

func (mainFlow *MainFlow) OnControllerKeyPress(controller model.Controller, key termbox.Key, ch rune) bool {
	if ch == 'q' {
		mainExecute := controller.GetEnvironment().GetValue("execute").(core.Execute)
		mainExecute.Shut()
	}

	return true
}
