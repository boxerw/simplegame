package core

type Container interface {
	AddChild(object Object) bool
	RemoveChild(entUID uint64)
	GetChild(entUID uint64) (Object, bool)
	GetChildList() []Object
}

type Child interface {
	GetParent() Object
}

type _Child interface {
	Child
	setParent(object Object)
}

type _ChildBase struct {
	parent Object
}

func (childBase *_ChildBase) shutChild(object Object) {
	if object.GetZombie() {
		return
	}

	if childBase.GetParent() != nil {
		childBase.GetParent().(Container).RemoveChild(object.GetUID())
	}
}

func (childBase *_ChildBase) GetParent() Object {
	return childBase.parent
}

func (childBase *_ChildBase) setParent(object Object) {
	childBase.parent = object
}
