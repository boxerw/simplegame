package core

func NewInput(ctx *Context, logic ...Logic) *Input {
	input := &Input{}
	input.initObject(ctx, input, logic...)
	return input
}

type Input struct {
	_ObjectBase
}

func (input *Input) Destroy() {
	input.shutObject()
}

func (input *Input) Update() {
	input.updateObject()
}
