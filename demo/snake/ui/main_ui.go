package ui

import (
	"simplegame/core"
	"simplegame/shell"
)

type MainUI struct {
	ani MainUIAni
}

func (ui *MainUI) Init(object core.Object, name string) {

}

func (ui *MainUI) Shut() {

}

func (ui *MainUI) Update(frameCtx core.FrameContext) {

}

type MainUIAni struct {
	randPosList []shell.Posi2D
}

func (ani *MainUIAni) Init(w, h, num int) {
	ani.randPosList = make([]shell.Posi2D, num)

	for i := 0; i < len(ani.randPosList); i++ {

	}
}

func (ani *MainUIAni) GetFrameMaps(frame int32) (shell.Maps, bool) {
	return nil, false
}

func (ani *MainUIAni) GetTotalFrames() int32 {
	return 0
}

func (ani *MainUIAni) GetFixedFPS() int32 {
	return 30
}
