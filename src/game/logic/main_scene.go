package logic

import (
	. "simple/game/core"
)

type MainScene struct {
	scene Object
}

func (logic *MainScene) Name() string {
	return "MainScene"
}

func (logic *MainScene) Init(object Object) {
	logic.scene = object
	logic.scene.GetContext().SetValue("MainScene", object)

	player := NewActor(logic.scene.GetContext(), &Player{})
	logic.scene.(Container).AddChild(player)
}

func (logic *MainScene) Shut() {
}

func (logic *MainScene) Update() {

}
