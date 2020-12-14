package core

func NewActor(ctx *Context, logic ...Logic) *Actor {
	actor := &Actor{}
	actor.initObject(ctx, actor, logic...)
	return actor
}

type Actor struct {
	_ObjectBase
	_ChildBase
	_TransformBase
}

func (actor *Actor) Destroy() {
	actor.shutChild(actor)
	actor.shutObject()
}

func (actor *Actor) Update() {
	actor.updateObject()
}
