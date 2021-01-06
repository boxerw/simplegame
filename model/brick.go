package model

import (
	"simplegame/core"
)

type Brick interface {
	core.Object
	core.Child
	core.Transform2D
}

func NewBrick(env core.Environment, components ...core.ComponentBundle) Brick {
	brick := &_Brick{}
	brick.ObjectModule.Init(env, brick, func() {
		brick.ChildModule.Init(brick)
	}, components...)
	return brick
}

type _Brick struct {
	core.ObjectModule
	core.ChildModule
	core.Transform2DModule
}

func (brick *_Brick) Destroy() {
	brick.ObjectModule.Shut(func() {
		brick.ChildModule.Shut(brick)
	})
}
