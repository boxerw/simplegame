package logic

import (
	"simplegame/core"
	"simplegame/shell"
)

type CoverStage struct {
	shell.Scene
}

func (cover *CoverStage) Init(object core.Object, name string) {
	cover.Scene = object.(shell.Scene)
}

func (cover *CoverStage) Shut() {
}

func (cover *CoverStage) Update(frameCtx core.FrameContext) {
}
