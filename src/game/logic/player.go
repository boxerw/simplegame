package logic

import (
	. "simple/game/core"
)

type LogicPlayer struct {
	actor Object
}

func (logic *LogicPlayer) Name() string {
	return "Player"
}

func (logic *LogicPlayer) Init(object Object) {
	logic.actor = object
}

func (logic *LogicPlayer) Shut() {
}

func (logic *LogicPlayer) Update() {
}
