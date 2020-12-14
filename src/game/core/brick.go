package core

func NewBrick(ctx *Context, logic ...Logic) *Brick {
	brick := &Brick{}
	brick.initObject(ctx, brick, logic...)
	return brick
}

type Brick struct {
	_ObjectBase
	_ChildBase
	_TransformBase
}

func (brick *Brick) Destroy() {
	brick.shutChild(brick)
	brick.shutObject()
}

func (brick *Brick) Update() {
	brick.updateObject()
}
