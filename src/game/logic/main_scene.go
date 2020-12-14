package logic

import (
	. "simple/game/core"
)

type MainScene struct {
	scene *Scene
}

func (logic *MainScene) Name() string {
	return "MainScene"
}

func (logic *MainScene) Init(object Object) {
	logic.scene = object.(*Scene)
	logic.scene.GetContext().SetValue("MainScene", logic.scene)

	player := NewActor(logic.scene.GetContext(), &Player{})
	logic.scene.AddChild(player)
}

func (logic *MainScene) Shut() {
}

func (logic *MainScene) Update() {
}
