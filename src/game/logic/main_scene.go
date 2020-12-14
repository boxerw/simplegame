package logic

import (
	. "simple/game/core"
)

type LogicMainScene struct {
	scene Object
}

func (logic *LogicMainScene) Name() string {
	return "MainScene"
}

func (logic *LogicMainScene) Init(object Object) {
	logic.scene = object
	logic.scene.GetContext().SetValue("MainScene", object)

	player := NewActor(logic.scene.GetContext(), &LogicPlayer{})
	logic.scene.(Container).AddChild(player)
}

func (logic *LogicMainScene) Shut() {
}

func (logic *LogicMainScene) Update() {

}
