package core

import "simplegame/core/foundation"

type Child = foundation.Child

type ChildModule struct {
	parent Object
}

func (childModule *ChildModule) Init(child Object) {
	childModule.parent = nil
}

func (childModule *ChildModule) Shut(child Object) {
	if childModule.GetParent() != nil {
		childModule.GetParent().(Container).RemoveChild(child.GetUID())
	}
}

func (childModule *ChildModule) GetParent() Object {
	return childModule.parent
}

func (childModule *ChildModule) SetParent(parent Object) {
	childModule.parent = parent
}
