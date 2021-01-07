package shell

import (
	"simplegame/core"
)

type Atom interface {
	core.Object
	core.Child
	core.Transform2D
	core.Data
}

func NewAtom(env core.Environment, components ...core.ComponentBundle) Atom {
	brick := &_Atom{}
	brick.ObjectModule.Init(env, brick, func() {
		brick.ChildModule.Init(brick)
	}, components...)
	return brick
}

type _Atom struct {
	core.ObjectModule
	core.ChildModule
	core.Transform2DModule
	core.DataModule
}

func (atom *_Atom) Destroy() {
	atom.ObjectModule.Shut(func() {
		atom.ChildModule.Shut(atom)
	})
}
