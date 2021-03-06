package logic

import (
	"github.com/boxerw/simplegame/client"
	"github.com/boxerw/simplegame/core"
	"github.com/nsf/termbox-go"
)

type MainFlowStage int

const (
	MainFlowStage_Init MainFlowStage = iota
	MainFlowStage_Cover
	MainFlowStage_Gaming
)

type MainFlow struct {
	client.Scene
	client.ControllerEventHook
	controller client.Controller
	screen     client.Screen
	stage      MainFlowStage
}

func (mainFlow *MainFlow) Init(object core.Object, name string) {
	scene, ok := object.(client.Scene)
	if !ok {
		panic("not scene")
	}

	mainFlow.Scene = scene

	mainFlow.controller = client.NewController(scene.GetEnvironment())
	mainFlow.controller.AddHook(mainFlow)

	mainFlow.screen = mainFlow.GetEnvironment().GetValue("screen").(client.Screen)
	mainFlow.stage = MainFlowStage_Init
}

func (mainFlow *MainFlow) Shut() {
	mainFlow.controller.Destroy()
}

func (mainFlow *MainFlow) Update(frameCtx core.FrameContext) {
	mainFlow.controller.Update(frameCtx)

	switch mainFlow.stage {
	case MainFlowStage_Init:
		mainFlow.SwitchStage(MainFlowStage_Cover)
	}
}

func (mainFlow *MainFlow) OnControllerKeyPress(controller client.Controller, key termbox.Key, ch rune) bool {
	switch mainFlow.stage {
	case MainFlowStage_Cover:
		if 'q' == ch {
			mainFlow.GetEnvironment().GetValue("execute").(core.Execute).Shut()
		} else {
			mainFlow.SwitchStage(MainFlowStage_Gaming)
		}
	case MainFlowStage_Gaming:
		if 'q' == ch {
			mainFlow.SwitchStage(MainFlowStage_Cover)
		}
	}

	return true
}

func (mainFlow *MainFlow) SwitchStage(newStage MainFlowStage) {
	if newStage == MainFlowStage_Init {
		return
	}

	switch newStage {
	case MainFlowStage_Cover:
		switch mainFlow.stage {
		case MainFlowStage_Init:
			mainFlow.AddComponent(core.NewComponentBundle("CoverStage", &CoverStage{}))

		case MainFlowStage_Gaming:
			mainFlow.RemoveComponent("GamingStage")
			mainFlow.AddComponent(core.NewComponentBundle("CoverStage", &CoverStage{}))
		}

	case MainFlowStage_Gaming:
		switch mainFlow.stage {
		case MainFlowStage_Cover:
			mainFlow.RemoveComponent("CoverStage")
			mainFlow.AddComponent(core.NewComponentBundle("GamingStage", &GamingStage{}))
		}
	}

	mainFlow.stage = newStage
}
