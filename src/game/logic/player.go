package logic

import (
	. "simple/game/core"
)

type Player struct {
	actor Object
}

func (logic *Player) Name() string {
	return "Player"
}

func (logic *Player) Init(object Object) {
	logic.actor = object
}

func (logic *Player) Shut() {
}

func (logic *Player) Update() {
}
