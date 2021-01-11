package logic

import (
	"github.com/nsf/termbox-go"
	"simplegame/core"
	"simplegame/shell"
)

type MainFlowStage int

const (
	MainFlowStage_Init MainFlowStage = iota
	MainFlowStage_Cover
	MainFlowStage_Gaming
)

type MainFlow struct {
	shell.Scene
	shell.ControllerEventHook
	controller shell.Controller
	screen     shell.Screen
	stage      MainFlowStage
}

func (mainFlow *MainFlow) Init(object core.Object, name string) {
	scene, ok := object.(shell.Scene)
	if !ok {
		panic("not scene")
	}

	mainFlow.Scene = scene

	mainFlow.controller = shell.NewController(scene.GetEnvironment())
	mainFlow.controller.AddHook(mainFlow)

	mainFlow.screen = mainFlow.GetEnvironment().GetValue("screen").(shell.Screen)
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

	case MainFlowStage_Cover:
		tips := "按'q'键退出，按其他任意键进入游戏"

		size := mainFlow.screen.GetCanvasSize()
		pos := shell.Posi2D{size.GetX()/2 - shell.StringWidth(tips)/2, int(float32(size.GetY()) * 0.8)}

		mainFlow.screen.DrawText(50, pos, tips, termbox.AttrBlink|termbox.ColorLightGray, termbox.ColorBlue)

	case MainFlowStage_Gaming:
		tips := "按'q'键返回"

		size := mainFlow.screen.GetCanvasSize()
		pos := shell.Posi2D{size.GetX()/2 - shell.StringWidth(tips)/2, int(float32(size.GetY()) * 0.8)}

		mainFlow.screen.DrawText(50, pos, tips, termbox.AttrBlink|termbox.ColorLightGray, termbox.ColorBlue)
	}
}

func (mainFlow *MainFlow) OnControllerKeyPress(controller shell.Controller, key termbox.Key, ch rune) bool {
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
