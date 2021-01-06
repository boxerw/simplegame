package model

import (
	"simple/core"
)

type Actor interface {
	core.Object
	core.Child
	core.Transform2D
	core.Data
}

func NewActor(env core.Environment, components ...core.ComponentBundle) Actor {
	actor := &_Actor{}
	actor.ObjectModule.Init(env, actor, func() {
		actor.ChildModule.Init(actor)
	}, components...)
	return actor
}

type _Actor struct {
	core.ObjectModule
	core.ChildModule
	core.Transform2DModule
	core.DataModule
}

func (actor *_Actor) Destroy() {
	actor.ObjectModule.Shut(func() {
		actor.ChildModule.Shut(actor)
	})
}
